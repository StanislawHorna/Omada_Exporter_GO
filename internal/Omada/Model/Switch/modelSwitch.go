package Switch

import (
	"omada_exporter_go/internal/Omada/Enum"
	"omada_exporter_go/internal/Omada/Model/Interface"
)

const path_OpenApiSwitch = "/openapi/v1/{omadaID}/sites/{siteID}/switches/{switchMac}"

// Implements Interface.Device
type Switch struct {
	DeviceType      Enum.DeviceType `json:"deviceType"`
	Name            string          `json:"name"`
	MacAddress      string          `json:"mac"`
	IP              string          `json:"ip"`
	IPv6List        []string        `json:"ipv6List"`
	Model           string          `json:"model"`
	FirmwareVersion string          `json:"firmwareVersion"`
	Version         string          `json:"version"`
	HardwareVersion string          `json:"hwVersion"`
	CpuUsage        int             `json:"cpuUtil"`
	RamUsage        int             `json:"memUtil"`
	Uptime          string          `json:"uptime"`
	PortList        []SwitchPort    `json:"portList"`
	LastSeen        float64
}

func (s Switch) GetType() string {
	return s.DeviceType.String()
}
func (s Switch) GetMacAddress() string {
	return s.MacAddress
}
func (s Switch) GetName() string {
	return s.Name
}
func (s Switch) GetIP() string {
	return s.IP
}
func (s Switch) GetModel() string {
	return s.Model
}
func (s Switch) GetHardwareVersion() string {
	return s.HardwareVersion
}
func (s Switch) GetFirmware() string {
	return s.FirmwareVersion
}
func (s Switch) GetCpuUsage() float64 {
	return float64(s.CpuUsage)
}
func (s Switch) GetMemUsage() float64 {
	return float64(s.RamUsage)
}
func (s Switch) GetTemperature() float64 {
	// Switches do not provide temperature data
	return Enum.NotApplicable_Float
}
func (s Switch) GetLastSeen() float64 {
	return s.LastSeen
}
func (s Switch) GetPorts() []Interface.Port {
	return Interface.ConvertToPortInterface(s.PortList)
}
