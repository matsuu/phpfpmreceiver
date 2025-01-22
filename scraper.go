package phpfpmreceiver

import (
	"context"
	"fmt"
	"time"

	"github.com/hipages/php-fpm_exporter/phpfpm"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/matsuu/phpfpmreceiver/internal/metadata"
)

type phpfpmScraper struct {
	pm *phpfpm.PoolManager

	settings component.TelemetrySettings
	cfg      *Config
	mb       *metadata.MetricsBuilder
}

func newPhpfpmScraper(settings receiver.Settings, cfg *Config) *phpfpmScraper {
	mb := metadata.NewMetricsBuilder(cfg.MetricsBuilderConfig, settings)
	return &phpfpmScraper{
		settings: settings.TelemetrySettings,
		cfg:      cfg,
		mb:       mb,
	}
}

func (r *phpfpmScraper) start(_ context.Context, _ component.Host) error {
	phpfpm.SetLogger(logrus.New())

	uris := r.cfg.Endpoint
	if len(uris) == 0 {
		return fmt.Errorf("no phpfpm scrape URI configured")
	}
	r.pm = new(phpfpm.PoolManager)
	for _, uri := range uris {
		r.pm.Add(uri)
	}

	return nil
}

func (r *phpfpmScraper) scrape(context.Context) (pmetric.Metrics, error) {
	if err := r.pm.Update(); err != nil {
		r.settings.Logger.Error("Failed to fetch phpfpm pool", zap.Error(err))
		return pmetric.Metrics{}, err
	}

	now := pcommon.NewTimestampFromTime(time.Now())
	for _, pool := range r.pm.Pools {
		r.mb.RecordPhpfpmAcceptedConnectionsDataPoint(now, pool.AcceptedConnections)
		r.mb.RecordPhpfpmActiveProcessesDataPoint(now, pool.ActiveProcesses)
		r.mb.RecordPhpfpmIdleProcessesDataPoint(now, pool.IdleProcesses)
		r.mb.RecordPhpfpmListenQueueDataPoint(now, pool.ListenQueue)
		r.mb.RecordPhpfpmMaxActiveProcessesDataPoint(now, pool.MaxActiveProcesses)
		r.mb.RecordPhpfpmMaxChildrenReachedDataPoint(now, pool.MaxChildrenReached)
		r.mb.RecordPhpfpmMaxListenQueueDataPoint(now, pool.MaxListenQueue)
		for i, process := range pool.Processes {
			child := int64(i)
			r.mb.RecordPhpfpmProcessLastRequestCPUDataPoint(now, process.LastRequestCPU, child)
			r.mb.RecordPhpfpmProcessLastRequestMemoryDataPoint(now, process.LastRequestMemory, child)
			r.mb.RecordPhpfpmProcessRequestsDataPoint(now, process.Requests, child)
			r.mb.RecordPhpfpmProcessStateDataPoint(now, process.Requests, child, process.State)
			r.mb.RecordPhpfpmProcessRequestDurationDataPoint(now, int64(process.RequestDuration), child)
		}
		r.mb.RecordPhpfpmScrapeFailuresDataPoint(now, pool.ScrapeFailures)
		r.mb.RecordPhpfpmSlowRequestsDataPoint(now, pool.SlowRequests)
		r.mb.RecordPhpfpmStartSinceDataPoint(now, pool.StartSince)
		rb := r.mb.NewResourceBuilder()
		rb.SetPhpfpmPoolName(pool.Name)
		rb.SetPhpfpmScrapeURI(pool.Address)
		r.mb.EmitForResource(metadata.WithResource(rb.Emit()))
	}

	return r.mb.Emit(), nil
}
