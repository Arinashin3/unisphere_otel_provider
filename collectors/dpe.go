package collectors

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"unisphere_otel/gounity/api"

	"go.opentelemetry.io/otel/metric"
)

func init() {
	key := "dpe"
	registerModule(key, NewDPE())
}

type ModuleDPE struct {
	// Module's Information
	name  string
	opts  *api.UnityActionOptions
	descs []*MetricDescriptor

	// Configuration File
	Enabled bool `yaml:"enabled"`
}

func NewDPE() *ModuleDPE {
	return &ModuleDPE{}
}

func (_m *ModuleDPE) Init(key string) {
	_m.Enabled = true
	_m.name = key

	_m.descs = []*MetricDescriptor{
		{
			Key:      "health.value",
			Name:     "unisphere_dpe_health",
			Desc:     "health about DPE of system",
			Unit:     "",
			TypeName: "gauge",
		},
		{
			Key:      "currentTemperature",
			Name:     "unisphere_dpe_current_temperature",
			Desc:     "current temperature of the DPE",
			Unit:     "",
			TypeName: "gauge",
		},
	}
	_m.opts = api.NewUnityActionOptions("dpe")
	_m.opts.Fields = []string{"id"}

	for _, desc := range _m.descs {
		_m.opts.Fields = append(_m.opts.Fields, desc.Key)
	}
}

// SetConfig
// allow module's config
func (_m *ModuleDPE) SetConfig(inf interface{}) Module {
	data, _ := json.Marshal(inf)
	json.NewDecoder(bytes.NewReader(data)).Decode(&_m)
	return _m
}

func (_m *ModuleDPE) Run(logger *slog.Logger, col *Collector) {
	meter := col.MeterProvider.Meter(_m.name)
	client := col.Client

	// Register Metrics...
	var observableMap map[string]metric.Float64Observable
	observableMap = CreateMapMetricDescriptor(meter, _m.descs, logger)

	// Register Metrics for Observables...
	var observableArray []metric.Observable
	for _, observable := range observableMap {
		observableArray = append(observableArray, observable)
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

		// Observer
		for _, v := range data {
			for observableKey, observable := range observableMap {
				observer.ObserveFloat64(observable, v.Get(observableKey).Float(), clientAttrs)
			}
		}

		return nil
	}, observableArray...)

}
