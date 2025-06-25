package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"omada_exporter_go/internal"
	"omada_exporter_go/internal/Log"
	"omada_exporter_go/internal/Prometheus"
)

func main() {
	Log.Init()
	Log.Info("Starting %s", internal.AppName)
	http.Handle("/omadaMetrics", Prometheus.OmadaMetricsHandler())
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
