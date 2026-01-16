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
	key := "basicSystemInfo"
	registerModule(key, NewBasicSystemInfo())
}

type ModuleBasicSystemInfo struct {
	// Module's Information
	name     string
	opts     *api.UnityActionOptions
	desc     []*MetricDescriptor
	defaults bool // Default Enabled

	// Configuration File
	Enabled *bool `yaml:"enabled"`
}

func NewBasicSystemInfo() *ModuleBasicSystemInfo {
	return &ModuleBasicSystemInfo{
		defaults: true,
	}
}

func (_m *ModuleBasicSystemInfo) Init(key string) {
	_m.name = key
	_m.desc = []*MetricDescriptor{
		{
			Key:      "info",
			Name:     "unisphere_basic_system_info",
			Desc:     "Information about unisphere basicSystemInfo",
			Unit:     "",
			TypeName: "gauge",
		},
	}
	_m.opts = api.NewUnityActionOptions("basicSystemInfo")
	_m.opts.Fields = []string{"model", "softwareFullVersion"}
}

func (_m *ModuleBasicSystemInfo) SetConfig(inf interface{}) Module {
	data, _ := json.Marshal(inf)
	json.NewDecoder(bytes.NewReader(data)).Decode(&_m)
	return _m
}

func (_m *ModuleBasicSystemInfo) Run(logger *slog.Logger, col *Collector) {
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
			infoAttrs := metric.WithAttributes(attribute.String("product.name", v.Get("model").String()), attribute.String("firmware.version", v.Get("softwareFullVersion").String()))
			observer.ObserveFloat64(observableMap["info"], 1, clientAttrs, infoAttrs)
		}

		return nil
	}, observableArray...)

}
