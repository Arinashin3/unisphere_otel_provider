package main

import (
	"context"
	"log/slog"
	"os"
	"unisphere_otel_provider/collectors"
	"unisphere_otel_provider/gounity"

	"github.com/Arinashin3/otel/config"

	"github.com/Arinashin3/otel/utils"
	"github.com/alecthomas/kingpin/v2"
	"github.com/prometheus/common/promslog"
	promslogflag "github.com/prometheus/common/promslog/flag"
)

const serviceName = "unisphere_otel_provider"

var (
	configFile = kingpin.Flag("config.file", "Paths to config file.").Short('c').Default("config.yml").String()
	logger     *slog.Logger
)

func main() {
	// Flag & Logger Configuration
	promslogConfig := &promslog.Config{}
	promslogflag.AddFlags(kingpin.CommandLine, promslogConfig)
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	logger = promslog.New(promslogConfig)

	// Load & Set Configuration
	cfg := config.NewConfiguration()
	cfg.LoadFile(*configFile, logger)

	var ctx = context.Background()
	mps := cfg.GenerateMeterProviders(ctx, serviceName)
	lps := cfg.GenerateLoggerProviders(ctx, serviceName)

	if !cfg.CheckSuccess() {
		logger.Error("failed to load config file")
		os.Exit(1)
	}

	trInsecure := utils.NewTransport(true)
	trSecure := utils.NewTransport(false)

	// Create Collectors... -> Clients
	var cols []*collectors.Collector
	for _, client := range cfg.Clients {
		col := collectors.NewCollector(ctx, client.Labels, *client.Interval)
		col.Instance = *client.Endpoint
		col.MeterProvider = mps[*client.Endpoint]
		col.LoggerProvider = lps[*client.Endpoint]

		// Create Clients...
		basicAuth := cfg.SearchBasicAuth(*client.Auth)
		switch *client.Insecure {
		case true:
			col.Client = gounity.NewUnisphereClient(*client.Endpoint, basicAuth, trInsecure)
		case false:
			col.Client = gounity.NewUnisphereClient(*client.Endpoint, basicAuth, trSecure)
		}
		cols = append(cols, col)

	}

	// Set Collector Configurations...
	for k, v := range cfg.Collectors {
		collectors.Modules[k].SetConfig(v)
	}

	// Run Collectors...
	for _, c := range cols {
		go c.Start(logger)
	}
	select {}

}
