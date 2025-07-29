package main

import (
	"net/http"

	"omada_exporter_go/internal"
	"omada_exporter_go/internal/Log"
	"omada_exporter_go/internal/Prometheus"
)

func main() {
	Log.Init()
	Log.Info("Starting %s", internal.AppName)
	http.Handle("/omadaMetrics", Prometheus.OmadaMetricsHandler())
	http.Handle("/metrics", Prometheus.OmadaMetricsHandler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
