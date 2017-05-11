package main

import (
	"log"
	"time"

	"github.com/jpillora/backoff"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	migrationEvent = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gce_instance_migration_event_active",
		Help: "Whether a GCE live migration event is pending; 1 if true, otherwise 0.",
	})
)

type GCEMetadataCollector struct {
	client *GCEMetadataClient
}

func NewGCEMetadataCollector() *GCEMetadataCollector {
	return &GCEMetadataCollector{
		client: NewGCEMetadataClient(),
	}
}

func (c *GCEMetadataCollector) Start() {
	go func() {
		boff := &backoff.Backoff{}

		var etag string
		for {
			var value string
			result, err := c.client.Get("instance/maintenance-event", &value, GetOptions{
				WaitForChange: true,
				LastETag:      etag,
			})
			if err != nil {
				log.Printf("Error getting metadata: %s", err)
				time.Sleep(boff.Duration())
				continue
			}

			switch value {
			case "MIGRATE_ON_HOST_MAINTENANCE":
				migrationEvent.Set(1)
			case "NONE":
				migrationEvent.Set(0)
			default:
				log.Printf("Warning: Received unexpected value %q from instance/maintenance-event endpoint", value)
			}

			etag = result.ETag
			boff.Reset()
		}
	}()
}

func init() {
	prometheus.MustRegister(migrationEvent)
}
