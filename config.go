package phpfpmreceiver

import (
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/matsuu/phpfpmreceiver/internal/metadata"
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	Endpoint                       []string `mapstructure:"endpoint"`
}
