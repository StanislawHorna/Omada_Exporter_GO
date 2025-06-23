package Prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"omada_exporter_go/internal/Omada/Model/Interface"
	"omada_exporter_go/internal/Prometheus/Utils"
)

const (
	label_portName     string = "portName"
	label_portIP       string = "portIP"
	label_portProtocol string = "portProtocol"
)

var portInfoLabels = []string{label_portName, label_portIP, label_portProtocol}

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
	port_speed = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_speed",
			Help: "Speed of the port in bits per second",
		},
		portIdentityLabels,
	)
	port_info = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "port_info",
			Help: "Information about port of the device",
		},
		append(portIdentityLabels, portInfoLabels...),
	)
)

func ExposePortMetrics(devices []Interface.Device) {
	for _, d := range devices {
		for _, p := range d.GetPorts() {

			labels := getPortIdentityLabels(d, p)
			port_rx_bytes_total.With(labels).Set(float64(p.GetRxBytes()))
			port_tx_bytes_total.With(labels).Set(float64(p.GetTxBytes()))

			port_speed.With(labels).Set(p.GetPortSpeed())
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
	},
	)).Set(1)
}
