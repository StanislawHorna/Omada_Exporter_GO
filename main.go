package main

import (
	"net/http"

	"omada_exporter_go/internal/Prometheus"
)

func main() {
	http.Handle("/metrics", Prometheus.MetricsHandler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
