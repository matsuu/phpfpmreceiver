package phpfpmreceiver

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/matsuu/phpfpmreceiver/internal/metadata"
)

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithMetrics(
			createMetricsReceiver,
			metadata.MetricsStability,
		),
	)
}

func createDefaultConfig() component.Config {
	cc := scraperhelper.NewDefaultControllerConfig()
	cc.CollectionInterval = 10 * time.Second
	mbc := metadata.DefaultMetricsBuilderConfig()

	defaultEndpoint := "tcp://127.0.0.1:9000/status"

	return &Config{
		ControllerConfig:     cc,
		MetricsBuilderConfig: mbc,
		Endpoint:             []string{defaultEndpoint},
	}
}

func createMetricsReceiver(_ context.Context, settings receiver.Settings, baseCfg component.Config, consumer consumer.Metrics) (receiver.Metrics, error) {
	cfg := baseCfg.(*Config)

	ps := newPhpfpmScraper(settings, cfg)
	s, err := scraper.NewMetrics(ps.scrape, scraper.WithStart(ps.start))
	if err != nil {
		return nil, err
	}

	return scraperhelper.NewMetricsController(
		&cfg.ControllerConfig,
		settings,
		consumer,
		scraperhelper.AddScraper(metadata.Type, s),
	)
}

var _ receiver.CreateMetricsFunc = createMetricsReceiver
