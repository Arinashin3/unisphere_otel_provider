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
	key := "ethernetPort"
	registerModule(key, NewEthernetPort())
}

func (_m *ModuleEthernetPort) SetConfig(inf interface{}) Module {
	data, _ := json.Marshal(inf)
	json.NewDecoder(bytes.NewReader(data)).Decode(&_m)
	return _m
}

type ModuleEthernetPort struct {
	// Module's Information
	name     string
	opts     *api.UnityActionOptions
	desc     []*MetricDescriptor
	defaults bool
	labels   []string

	// Configuration File
	Enabled *bool
}

func NewEthernetPort() *ModuleEthernetPort {
	return &ModuleEthernetPort{
		defaults: false,
	}
}

func (_m *ModuleEthernetPort) Init(key string) {
	_m.name = key
	_m.desc = []*MetricDescriptor{
		{
			Key:      "info",
			Name:     "unisphere_ethernetPort_info",
			Desc:     "Information of the associated resource",
			Unit:     "",
			TypeName: "gauge",
		},
		{
			Key:      "health.value",
			Name:     "unisphere_ethernetPort_health",
			Desc:     "Health of the associated resource",
			Unit:     "",
			TypeName: "gauge",
		},
		{
			Key:      "speed",
			Name:     "unisphere_ethernetPort_speed",
			Desc:     "Supported Ethernet port transmission speeds",
			Unit:     "mbps",
			TypeName: "gauge",
		},
		{
			Key:      "isLinkUp",
			Name:     "unisphere_ethernetPort_is_link_up",
			Desc:     "Indicates whether the Ethernet port has link. (Applies if the Ethernet port is configured with a link.) Indicates whether the Ethernet port's link is up",
			Unit:     "",
			TypeName: "gauge",
		},
	}
	_m.labels = []string{"id", "name"}
	_m.opts = api.NewUnityActionOptions("ethernetPort")
	for _, v := range _m.desc {
		if v.Key == "info" {
			continue
		}
		_m.opts.Fields = append(_m.opts.Fields, v.Key)
	}
	_m.opts.Fields = append(_m.opts.Fields, _m.labels...)
}

func (_m *ModuleEthernetPort) Run(logger *slog.Logger, col *Collector) {
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
			return nil
		}
		clientAttrs := metric.WithAttributes(append(col.customLabels, col.detectLabels...)...)

		// Request Data
		data, err := client.GetInstances(_m.opts)
		if err != nil {
			logger.Error("Failed to get", "error", err, "module", _m.name)
			col.success = false
			return nil
		}
		col.success = true

		// Capacity Attributes...
		for _, v := range data {
			ethernetPortAttrs := metric.WithAttributes(
				attribute.String("fePort", v.Get("id").String()),
			)
			infoAttrs := metric.WithAttributes(
				attribute.String("fePort.name", v.Get("name").String()),
			)
			for _, desc := range _m.desc {
				key := desc.Key
				var f float64
				switch key {
				case "info":
					observer.ObserveFloat64(observableMap[key], 1, clientAttrs, ethernetPortAttrs, infoAttrs)
					continue
				case "isLinkUp":
					if v.Get(key).Bool() {
						f = 1
					} else {
						f = 0
					}
				default:
					f = v.Get(key).Float()
				}
				observer.ObserveFloat64(observableMap[key], f, clientAttrs, ethernetPortAttrs)
			}
		}

		return nil
	}, observableArray...)

}
