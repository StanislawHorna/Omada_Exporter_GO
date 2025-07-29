package Prometheus

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"omada_exporter_go/internal/Omada/Enum"
	"omada_exporter_go/internal/Omada/Model/Interface"
	"omada_exporter_go/internal/Prometheus/Utils"
)

const (
	label_portName     string = "portName"
	label_portIP       string = "portIP"
	label_portProtocol string = "portProtocol"
	label_portMode     string = "portMode"
)

var portInfoLabels = []string{label_portName, label_portIP, label_portProtocol, label_portMode}

var (
	port_rx_bytes_total = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_rx_bytes_total",
			Help: "Total number of bytes received on the port",
		},
		portIdentityLabels,
	)
	port_tx_bytes_total = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_tx_bytes_total",
			Help: "Total number of bytes transmitted on the port",
		},
		portIdentityLabels,
	)
	port_speed = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_speed",
			Help: "Speed of the port in bits per second",
		},
		portIdentityLabels,
	)
	port_duplex = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_duplex",
			Help: fmt.Sprintf(
				"Duplex mode of the port (%s)",
				Enum.GetDuplexPossibleValues(),
			),
		},
		portIdentityLabels,
	)
	port_info = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_info",
			Help: "Information about port of the device",
		},
		append(portIdentityLabels, portInfoLabels...),
	)
	port_upstream_state = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_upstream_state",
			Help: fmt.Sprintf("Router port upstream state (%s)",
				Enum.GetRouterUpstreamStatePossibleValues(),
			),
		},
		portIdentityLabels,
	)
	port_internet_state = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_internet_state",
			Help: fmt.Sprintf("Router port Internet state (%s)",
				Enum.GetInternetStatePossibleValues(),
			),
		},
		portIdentityLabels,
	)
	port_internet_latency = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_internet_latency",
			Help: "Router port Internet latency in milliseconds",
		},
		portIdentityLabels,
	)
	port_internet_loss = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_internet_loss",
			Help: "Router port Internet loss in percentage (0 - 100)",
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

			port_speed.With(labels).Set(p.GetPortSpeed())
			port_duplex.With(labels).Set(p.GetPortDuplex())

			if p.GetInternetState() != Enum.NotApplicable_Float {
				port_internet_state.With(labels).Set(p.GetInternetState())
			}
			if p.GetUpstreamState() != Enum.NotApplicable_Float {
				port_upstream_state.With(labels).Set(p.GetUpstreamState())
			}
			if p.GetInternetLatency() != Enum.NotApplicable_Float {
				port_internet_latency.With(labels).Set(p.GetInternetLatency())
			}
			if p.GetInternetLoss() != Enum.NotApplicable_Float {
				port_internet_loss.With(labels).Set(p.GetInternetLoss())
			}

			setPortInfo(p, labels)
		}
	}

}

func setPortInfo(port Interface.Port, labels prometheus.Labels) {
	// Delete all info metrics to avoid duplicates created due to changed labels
	// new set of labels always creates new series, but old one is not deleted,
	// even if it was not set in the current iteration
	port_info.Delete(labels)
	port_info.With(Utils.AppendMaps(labels, map[string]string{
		label_portName:     port.GetPortName(),
		label_portIP:       port.GetPortIP(),
		label_portProtocol: port.GetPortProtocol(),
		label_portMode:     port.GetPortMode(),
	},
	)).Set(1)
}
