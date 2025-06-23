package Prometheus

import (
	"github.com/prometheus/client_golang/prometheus"

	"omada_exporter_go/internal/Omada/Model/Interface"
)

const (
	label_deviceType string = "deviceType"
	label_macAddress string = "macAddress"
	label_portID     string = "portID"
)

var deviceIdentityLabels = []string{label_deviceType, label_macAddress}

var portIdentityLabels = []string{label_deviceType, label_macAddress, label_portID}

var omadaRegistry = prometheus.NewRegistry()

func getDeviceIdentityLabels(device Interface.Device) prometheus.Labels {
	return prometheus.Labels{
		label_deviceType: device.GetType(),
		label_macAddress: device.GetMacAddress(),
	}
}

func getPortIdentityLabels(device Interface.Device, port Interface.Port) prometheus.Labels {
	return prometheus.Labels{
		label_deviceType: device.GetType(),
		label_macAddress: device.GetMacAddress(),
		label_portID:     port.GetID(),
	}
}
