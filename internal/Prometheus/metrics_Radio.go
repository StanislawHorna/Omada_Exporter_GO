package Prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"omada_exporter_go/internal/Omada/Model/Interface"
)

var (
	radio_tx_bytes_total = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_tx_bytes_total",
			Help: "Total number of bytes transmitted on the radio",
		},
		radioIdentityLabels,
	)
	radio_rx_bytes_total = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_rx_bytes_total",
			Help: "Total number of bytes received on the radio",
		},
		radioIdentityLabels,
	)
)

func ExposeRadioMetrics(devices []Interface.Device) {
	for _, d := range devices {
		if d.GetRadios() == nil {
			continue
		}
		for _, r := range d.GetRadios() {
			labels := getRadioIdentityLabels(d, r)

			radio_rx_bytes_total.With(labels).Set(r.GetRxBytes())
			radio_tx_bytes_total.With(labels).Set(r.GetTxBytes())
		}
	}
}
