package Prometheus

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"omada_exporter_go/internal/Omada/Model/AccessPoint"
	"omada_exporter_go/internal/Omada/Model/Devices"
	"omada_exporter_go/internal/Omada/Model/Gateway"
	"omada_exporter_go/internal/Omada/Model/Interface"
	"omada_exporter_go/internal/Omada/Model/Switch"
)

var (
	omada_scrape_duration = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "omada_scrape_duration",
			Help: "The amount of time it took to scrape the Omada controller in seconds",
		},
	)
)

func CollectMetrics() error {
	start := time.Now()
	deviceList, err := Devices.Get()
	if err != nil {
		return err
	}

	var omadaDevices []Interface.Device

	switches, err := Switch.Get(*deviceList)
	if err == nil {
		for i := range *switches {
			(*switches)[i].AppendGeneralProperties(deviceList)
		}
		Interface.AppendDevicesSlice(&omadaDevices, *switches)
	} else {
		fmt.Println("failed to get switches: %w", err)
	}

	gateways, err := Gateway.Get(*deviceList)
	if err == nil {
		for i := range *gateways {
			(*gateways)[i].AppendGeneralProperties(deviceList)
		}
		Interface.AppendDevicesSlice(&omadaDevices, *gateways)
	} else {
		fmt.Println("failed to get gateways: %w", err)
	}

	aps, err := AccessPoint.Get(*deviceList)
	if err == nil {
		for i := range *aps {
			(*aps)[i].AppendGeneralProperties(deviceList)
		}
		Interface.AppendDevicesSlice(&omadaDevices, *aps)
	} else {
		fmt.Println("failed to get access points: %w", err)
	}
	omada_scrape_duration.Set(time.Since(start).Seconds())

	ExposeDeviceMetrics(omadaDevices)
	ExposePortMetrics(omadaDevices)
	ExposeRadioMetrics(omadaDevices)
	return err
}
