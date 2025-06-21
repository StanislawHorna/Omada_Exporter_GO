package Prometheus

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func OmadaMetricsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		CollectMetrics()
		promhttp.HandlerFor(omadaRegistry, promhttp.HandlerOpts{}).ServeHTTP(w, r)
	})
}
