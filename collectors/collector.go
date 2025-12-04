package collectors

import (
	"context"
	"errors"
	"log/slog"
	"time"
	"unisphere_otel/gounity"
	"unisphere_otel/gounity/api"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	sdkLog "go.opentelemetry.io/otel/sdk/log"
	sdkMetric "go.opentelemetry.io/otel/sdk/metric"
)

var Modules = make(map[string]Module)

type Module interface {
	Run(logger *slog.Logger, col *Collector)
	Init(key string)
	SetConfig(inf interface{}) Module
}

func registerModule(name string, module Module) {
	module.Init(name)
	Modules[name] = module
}

type Collector struct {
	ctx            context.Context
	Instance       string
	customLabels   []attribute.KeyValue
	detectLabels   []attribute.KeyValue
	MeterProvider  *sdkMetric.MeterProvider
	LoggerProvider *sdkLog.LoggerProvider
	interval       time.Duration
	Client         *gounity.UnisphereClient
	success        bool
}

func NewCollector(ctx context.Context, attrs map[string]string, interval time.Duration) *Collector {
	var customLabels []attribute.KeyValue
	for k, v := range attrs {
		customLabels = append(customLabels, attribute.String(k, v))
	}
	return &Collector{
		ctx:          ctx,
		customLabels: customLabels,
		interval:     interval,
	}
}

func (_col *Collector) Start(logger *slog.Logger) {
	// Init HostName
	opt := api.NewUnityActionOptions("system")
	opt.Fields = []string{"name"}
	data, err := _col.Client.GetInstances(opt)
	if err != nil {
		logger.Warn("cannot set labels", "error", err)
	} else {
		for _, v := range data {
			_col.detectLabels = append(_col.detectLabels, attribute.String("host.name", v.Get("name").String()))
		}

	}

	for _, v := range Modules {
		go v.Run(logger, _col)
	}
	select {}
}

type MetricDescriptor struct {
	Key      string
	Name     string
	Desc     string
	Unit     string
	TypeName string
}

func CreateMapMetricDescriptor(meter metric.Meter, mds []*MetricDescriptor, logger *slog.Logger) map[string]metric.Float64Observable {
	mdmap := make(map[string]metric.Float64Observable)
	var err error
	for _, md := range mds {
		var tmp metric.Float64Observable
		desc := metric.WithDescription(md.Desc)
		unit := metric.WithUnit(md.Unit)
		switch md.TypeName {
		case "counter":
			tmp, err = meter.Float64ObservableCounter(md.Name, desc, unit)
		case "gauge":
			tmp, err = meter.Float64ObservableGauge(md.Name, desc, unit)
		default:
			err = errors.New("unknown metric type")
		}
		if err != nil {
			logger.Warn("cannot create metric", "error", err, "metric_key", md.Key, "metric_type", md.TypeName)
		}
		mdmap[md.Key] = tmp
	}
	return mdmap

}
