package Prometheus

import (
	"omada_exporter_go/internal/Omada/Model"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	label_deviceType string = "deviceType"
	label_macAddress string = "macAddress"
)

var identityLabels = []string{label_deviceType, label_macAddress}

var omadaRegistry = prometheus.NewRegistry()

func getIdentityLabels(device Model.DeviceInterface) prometheus.Labels {
	return prometheus.Labels{
		label_deviceType: device.GetType(),
		label_macAddress: device.GetMacAddress(),
	}
}
