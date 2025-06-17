package Prometheus

import (
	"github.com/prometheus/client_golang/prometheus"

	"omada_exporter_go/internal/Omada/Model/AccessPoint"
	"omada_exporter_go/internal/Omada/Model/Devices"
	"omada_exporter_go/internal/Omada/Model/Gateway"
	"omada_exporter_go/internal/Omada/Model/Switch"
)

func CollectMetrics() error {
	deviceList, err := Devices.Get()
	if err != nil {
		return err
	}

	switches, err := Switch.Get(*deviceList)
	if err != nil {
		return err
	}
	for _, s := range *switches {
		cpuUsage.With(
			prometheus.Labels{
				"deviceType": string(s.DeviceType),
				"deviceName": s.Name,
				"macAddress": s.MacAddress,
			},
		).Set(float64(s.CpuUsage))
	}

	gateways, err := Gateway.Get(*deviceList)
	if err != nil {
		return err
	}
	for _, g := range *gateways {
		cpuUsage.With(
			prometheus.Labels{
				"deviceType": string(g.DeviceType),
				"deviceName": g.Name,
				"macAddress": g.MacAddress,
			},
		).Set(float64(g.CpuUsage))
	}

	aps, err := AccessPoint.Get(*deviceList)
	if err != nil {
		return err
	}
	for _, ap := range *aps {
		cpuUsage.With(
			prometheus.Labels{
				"deviceType": string(ap.DeviceType),
				"deviceName": ap.Name,
				"macAddress": ap.MacAddress,
			},
		).Set(float64(ap.CpuUsage))
	}

	return nil
}
