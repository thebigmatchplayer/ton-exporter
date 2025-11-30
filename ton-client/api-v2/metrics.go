package apiv2

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	v2_BlockHeight prometheus.Gauge
}

func InitMetrics() *Metrics {
	m := &Metrics{
		v2_BlockHeight: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "ton_v2_BlockHeight",
			Help: "Current seq no of TON",
		}),
	}

	prometheus.MustRegister(m.v2_BlockHeight)
	return m
}
