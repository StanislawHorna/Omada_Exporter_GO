package Prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	cpuUsage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage",
			Help: "The percentage of CPU usage",
		},
		[]string{
			"deviceType",
			"deviceName",
			"macAddress",
		},
	)
)
