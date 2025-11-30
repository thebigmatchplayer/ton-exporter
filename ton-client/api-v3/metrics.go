package apiv3

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	v3_BlockHeight prometheus.Gauge
}

func InitMetrics() *Metrics {
	m := &Metrics{
		v3_BlockHeight: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "ton_v3_BlockHeight",
			Help: "Current seq no of TON",
		}),
	}

	prometheus.MustRegister(m.v3_BlockHeight)
	return m
}
