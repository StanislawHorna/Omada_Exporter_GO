package Prometheus

import (
	"fmt"

	"omada_exporter_go/internal/Omada/Model/AccessPoint"
	"omada_exporter_go/internal/Omada/Model/Devices"
	"omada_exporter_go/internal/Omada/Model/Gateway"
	"omada_exporter_go/internal/Omada/Model/Interface"
	"omada_exporter_go/internal/Omada/Model/Switch"
)

func CollectMetrics() error {
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

	ExposeDeviceMetrics(omadaDevices)
	ExposePortMetrics(omadaDevices)
	ExposeRadioMetrics(omadaDevices)
	return nil
}
