package Model

import "fmt"

type DeviceInterface interface {
	// Getters for device properties
	GetType() string
	GetMacAddress() string
	GetName() string
	GetIP() string
	GetFirmware() string
	GetModel() string

	// Getters for device resource consumption
	GetCpuUsage() float64
	GetMemUsage() float64

	GetTemperature() float64 // Returns -1 if temperature is not available
}

func AppendDevicesSlice[T DeviceInterface](devices *[]DeviceInterface, newDevices []T) error {
	if devices == nil {
		return fmt.Errorf("devices is nil")
	}

	for i := range newDevices {
		*devices = append(*devices, newDevices[i])
	}

	return nil
}
