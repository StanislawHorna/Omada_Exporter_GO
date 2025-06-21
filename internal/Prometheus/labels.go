package Prometheus

import (
	"github.com/prometheus/client_golang/prometheus"

	"omada_exporter_go/internal/Omada/Model/Interface"
)

const (
	label_deviceType string = "deviceType"
	label_macAddress string = "macAddress"
)

var identityLabels = []string{label_deviceType, label_macAddress}

var omadaRegistry = prometheus.NewRegistry()

func getIdentityLabels(device Interface.Device) prometheus.Labels {
	return prometheus.Labels{
		label_deviceType: device.GetType(),
		label_macAddress: device.GetMacAddress(),
	}
}
