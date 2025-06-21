package Devices

import "omada_exporter_go/internal/Omada/Enum"

const path_OpenApiDevicesList = "/openapi/v1/{omadaID}/sites/{siteID}/devices"

type DeviceStatus int

const (
	DeviceStatus_disconnected     DeviceStatus = 0
	DeviceStatus_connected        DeviceStatus = 1
	DeviceStatus_pending          DeviceStatus = 2
	DeviceStatus_heartbeatMissing DeviceStatus = 3
	DeviceStatus_isolated         DeviceStatus = 4
)

type Device struct {
	MacAddress string          `json:"mac"`
	Name       string          `json:"name"`
	Type       Enum.DeviceType `json:"type"`
	Model      string          `json:"model"`
	IP         string          `json:"ip"`
	Uptime     string          `json:"uptime"`
	LastSeen   float64         `json:"lastSeen"`
	Status     DeviceStatus    `json:"status"`
	CpuUsage   int             `json:"cpuUtil"`
	RamUsage   int             `json:"memUtil"`
	TagName    string          `json:"tagName"`
}

func (d *Device) GetStatus() string {
	switch d.Status {
	case DeviceStatus_disconnected:
		return "Disconnected"
	case DeviceStatus_connected:
		return "Connected"
	case DeviceStatus_pending:
		return "Pending"
	case DeviceStatus_heartbeatMissing:
		return "Heartbeat Missing"
	case DeviceStatus_isolated:
		return "Isolated"
	default:
		return "Unknown"
	}
}
