package Prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"omada_exporter_go/internal/Omada/Model/Interface"
	"omada_exporter_go/internal/Prometheus/Utils"
)

const (
	label_channel   string = "channel"
	label_bandwidth string = "bandwidth"
	label_mode      string = "mode"
)

var radioInfoLabels = []string{label_channel, label_bandwidth, label_mode}

var (
	radio_tx_bytes_total = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_tx_bytes_total",
			Help: "Total number of bytes transmitted on the radio",
		},
		radioIdentityLabels,
	)
	radio_rx_bytes_total = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_rx_bytes_total",
			Help: "Total number of bytes received on the radio",
		},
		radioIdentityLabels,
	)

	radio_rx_drop_packets_total = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_rx_drop_packets_total",
			Help: "Total number of Rx packets dropped on the radio",
		},
		radioIdentityLabels,
	)
	radio_tx_drop_packets_total = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_tx_drop_packets_total",
			Help: "Total number of Tx packets dropped on the radio",
		},
		radioIdentityLabels,
	)
	radio_rx_err_packets_total = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_rx_err_packets_total",
			Help: "Total number of Rx packets with error on the radio",
		},
		radioIdentityLabels,
	)
	radio_tx_err_packets_total = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_tx_err_packets_total",
			Help: "Total number of Tx packets with errors on the radio",
		},
		radioIdentityLabels,
	)
	radio_rx_retry_packets_total = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_rx_retry_packets_total",
			Help: "Total number of Rx packets retried on the radio",
		},
		radioIdentityLabels,
	)
	radio_tx_retry_packets_total = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_tx_retry_packets_total",
			Help: "Total number of Tx packets retried on the radio",
		},
		radioIdentityLabels,
	)

	radio_tx_usage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_tx_usage",
			Help: "Radio TX channel usage in percentage (0 - 100)",
		},
		radioIdentityLabels,
	)
	radio_rx_usage = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_rx_usage",
			Help: "Radio RX channel usage in percentage (0 - 100)",
		},
		radioIdentityLabels,
	)
	radio_interference = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_interference",
			Help: "Information about radio interface of the device",
		},
		radioIdentityLabels,
	)
	radio_max_tx_rate = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_max_tx_rate",
			Help: "Maximum transmission rate of the radio in bits per second",
		},
		radioIdentityLabels,
	)
	radio_info = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "radio_info",
			Help: "Information about radio interface of the device",
		},
		append(radioIdentityLabels, radioInfoLabels...),
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
			radio_interference.With(labels).Set(r.GetInterference())

			radio_max_tx_rate.With(labels).Set(r.GetMaxTxRate())

			setRadioInfo(r, labels)
		}
	}
}

func setRadioInfo(radio Interface.Radio, labels prometheus.Labels) {
	// Delete all info metrics to avoid duplicates created due to changed labels
	// new set of labels always creates new series, but old one is not deleted,
	// even if it was not set in the current iteration
	radio_info.Delete(labels)
	radio_info.With(Utils.AppendMaps(labels, map[string]string{
		label_channel:   radio.GetActualChannel(),
		label_bandwidth: radio.GetBandwidth(),
		label_mode:      radio.GetMode(),
	})).Set(1)
}
