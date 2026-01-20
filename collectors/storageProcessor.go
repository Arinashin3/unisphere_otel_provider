package collectors

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"unisphere_otel_provider/gounity/api"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

func init() {
	key := "storageProcessor"
	registerModule(key, NewStorageProcessor())
}

type ModuleStorageProcessor struct {
	// Module's Information
	name     string
	opts     *api.UnityActionOptions
	desc     []*MetricDescriptor
	defaults bool // Default Enabled

	// Configuration File
	Enabled *bool `yaml:"enabled"`
}

func NewStorageProcessor() *ModuleStorageProcessor {
	return &ModuleStorageProcessor{
		defaults: false,
	}
}

func (_m *ModuleStorageProcessor) Init(key string) {
	_m.name = key
	_m.desc = []*MetricDescriptor{
		{
			Key:      "info",
			Name:     "unisphere_storage_processor_info",
			Desc:     "Information about unisphere storage processor",
			Unit:     "",
			TypeName: "gauge",
		},
		{
			Key:      "health.Value",
			Name:     "unisphere_storage_processor_health",
			Desc:     "Health of unisphere storage processor",
			Unit:     "",
			TypeName: "gauge",
		},
		{
			Key:      "memorySize",
			Name:     "unisphere_storage_processor_memory_size",
			Desc:     "Memory Size of unisphere storage processor",
			Unit:     "mb",
			TypeName: "gauge",
		},
	}
	_m.opts = api.NewUnityActionOptions("storageProcessor")
	_m.opts.Fields = []string{"model", "id"}
	for _, m := range _m.desc {
		if m.Key == "info" {
			continue
		}
		_m.opts.Fields = append(_m.opts.Fields, m.Key)
	}
}

func (_m *ModuleStorageProcessor) SetConfig(inf interface{}) Module {
	data, _ := json.Marshal(inf)
	json.NewDecoder(bytes.NewReader(data)).Decode(&_m)
	return _m
}

func (_m *ModuleStorageProcessor) Run(logger *slog.Logger, col *Collector) {
	meter := col.MeterProvider.Meter(_m.name)
	client := col.Client

	// Register Metrics...
	var observableMap map[string]metric.Float64Observable
	observableMap = CreateMapMetricDescriptor(meter, _m.desc, logger)

	// Register Metrics for Observables...
	var observableArray []metric.Observable
	for _, obserable := range observableMap {
		observableArray = append(observableArray, obserable)
	}

	// Callback
	meter.RegisterCallback(func(ctx context.Context, observer metric.Observer) error {

		// Set Attributes
		if col.detectLabels == nil {
			logger.Debug("hostLabels not set")
			return nil
		}
		clientAttrs := metric.WithAttributes(append(col.customLabels, col.detectLabels...)...)

		// Request Data
		data, err := client.GetInstances(_m.opts)
		if err != nil {
			logger.Error("Failed to get", "error", err, "module", _m.name)
			return nil
		}

		// System Attributes...
		for _, v := range data {
			spAttrs := metric.WithAttributes(
				attribute.String("sp.id", v.Get("id").String()),
			)
			infoAttrs := metric.WithAttributes(attribute.String("sp.model", v.Get("model").String()))
			observer.ObserveFloat64(observableMap["info"], 1, clientAttrs, infoAttrs, spAttrs)
			observer.ObserveFloat64(observableMap["health.Value"], v.Get("health.Value").Float(), clientAttrs, spAttrs)
			observer.ObserveFloat64(observableMap["memorySize"], v.Get("memorySize").Float(), clientAttrs, spAttrs)
		}

		return nil
	}, observableArray...)

}
