package Prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"omada_exporter_go/internal/Omada/Model/Interface"
)

var (
	port_rx_bytes_total = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_rx_bytes_total",
			Help: "Total number of bytes received on the port",
		},
		portIdentityLabels,
	)
	port_tx_bytes_total = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_tx_bytes_total",
			Help: "Total number of bytes transmitted on the port",
		},
		portIdentityLabels,
	)
)

func ExposePortMetrics(devices []Interface.Device) {
	for _, d := range devices {
		for _, p := range d.GetPorts() {

			labels := getPortIdentityLabels(d, p)
			port_rx_bytes_total.With(labels).Set(float64(p.GetRxBytes()))
			port_tx_bytes_total.With(labels).Set(float64(p.GetTxBytes()))
		}
	}

}
