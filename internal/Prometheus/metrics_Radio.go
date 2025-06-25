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

	radio_rx_drop_packets_total = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_rx_drop_packets_total",
			Help: "Total number of Rx packets dropped on the radio",
		},
		radioIdentityLabels,
	)
	radio_tx_drop_packets_total = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_tx_drop_packets_total",
			Help: "Total number of Tx packets dropped on the radio",
		},
		radioIdentityLabels,
	)
	radio_rx_err_packets_total = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_rx_err_packets_total",
			Help: "Total number of Rx packets with error on the radio",
		},
		radioIdentityLabels,
	)
	radio_tx_err_packets_total = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_tx_err_packets_total",
			Help: "Total number of Tx packets with errors on the radio",
		},
		radioIdentityLabels,
	)
	radio_rx_retry_packets_total = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_rx_retry_packets_total",
			Help: "Total number of Rx packets retried on the radio",
		},
		radioIdentityLabels,
	)
	radio_tx_retry_packets_total = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_tx_retry_packets_total",
			Help: "Total number of Tx packets retried on the radio",
		},
		radioIdentityLabels,
	)

	radio_tx_usage = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_tx_usage",
			Help: "Radio TX channel usage in percentage (0 - 100)",
		},
		radioIdentityLabels,
	)
	radio_rx_usage = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_rx_usage",
			Help: "Radio RX channel usage in percentage (0 - 100)",
		},
		radioIdentityLabels,
	)
	radio_interface = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_interface",
			Help: "Information about radio interface of the device",
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

			radio_tx_bytes_total.With(labels).Set(r.GetTxBytes())
			radio_rx_bytes_total.With(labels).Set(r.GetRxBytes())

			radio_tx_drop_packets_total.With(labels).Set(r.GetTxDrops())
			radio_rx_drop_packets_total.With(labels).Set(r.GetRxDrops())
			radio_tx_err_packets_total.With(labels).Set(r.GetTxErrors())
			radio_rx_err_packets_total.With(labels).Set(r.GetRxErrors())
			radio_tx_retry_packets_total.With(labels).Set(r.GetTxRetries())
			radio_rx_retry_packets_total.With(labels).Set(r.GetRxRetries())

			radio_tx_usage.With(labels).Set(r.GetTxUsage())
			radio_rx_usage.With(labels).Set(r.GetRxUsage())
			radio_interface.With(labels).Set(r.GetInterference())
		}
	}
}
