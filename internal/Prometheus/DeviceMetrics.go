package Prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"omada_exporter_go/internal/Omada/Model/Interface"
	"omada_exporter_go/internal/Prometheus/Utils"
)

const (
	label_deviceName     string = "deviceName"
	label_deviceModel    string = "deviceModel"
	label_IP             string = "IP"
	label_deviceFirmware string = "deviceFirmware"
)

var (
	cpu_usage = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage",
			Help: "The percentage of CPU usage",
		},
		deviceIdentityLabels,
	)

	memory_usage = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "memory_usage",
			Help: "The percentage of memory usage",
		},
		deviceIdentityLabels,
	)

	temperature = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "temperature",
			Help: "The device temperature in degrees Celsius",
		},
		deviceIdentityLabels,
	)

	device_info = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "device_info",
			Help: "Information about the device",
		},
		append(deviceIdentityLabels, []string{label_deviceName, label_deviceModel, label_IP, label_deviceFirmware}...),
	)

	device_last_seen = promauto.With(omadaRegistry).NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "device_last_seen",
			Help: "The last time the device was seen, in Unix timestamp format",
		},
		deviceIdentityLabels,
	)
)

func ExposeDeviceMetrics(devices []Interface.Device) {
	for _, d := range devices {
		identityLabels := getDeviceIdentityLabels(d)

		cpu_usage.With(identityLabels).Set(d.GetCpuUsage())
		memory_usage.With(identityLabels).Set(d.GetMemUsage())
		device_last_seen.With(identityLabels).Set(d.GetLastSeen())

		setDeviceTemperature(d, identityLabels)
		setDeviceInfo(d, identityLabels)
	}

}

func setDeviceTemperature(device Interface.Device, labels prometheus.Labels) {
	temp := device.GetTemperature()
	if temp >= 0 {
		temperature.With(labels).Set(temp)
	} else {
		temperature.Delete(labels)
	}
}

func setDeviceInfo(device Interface.Device, labels prometheus.Labels) {
	// Delete all info metrics to avoid duplicates created due to changed labels
	// new set of labels always creates new series, but old one is not deleted,
	// even if it was not set in the current iteration
	device_info.DeletePartialMatch(labels)
	device_info.With(Utils.AppendMaps(map[string]string{
		label_deviceName:     device.GetName(),
		label_deviceModel:    device.GetModel(),
		label_IP:             device.GetIP(),
		label_deviceFirmware: device.GetFirmware(),
	}, labels)).Set(1)
}
