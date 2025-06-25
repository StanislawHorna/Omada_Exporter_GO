package Gateway

import (
	"omada_exporter_go/internal/Omada/Enum"
	"omada_exporter_go/internal/Omada/Model/Interface"
)

// Implements Interface.Device
type Gateway struct {
	// OpenAPI fields
	DeviceType      Enum.DeviceType `json:"deviceType"`
	Name            string          `json:"name"`
	MacAddress      string          `json:"mac"`
	Model           string          `json:"showModel"`
	FirmwareVersion string          `json:"firmwareVersion"`
	IP              string          `json:"ip"`
	Uptime          string          `json:"uptime"`
	Temperature     int             `json:"temp"`
	CpuUsage        int             `json:"cpuUtil"`
	RamUsage        int             `json:"memUtil"`
	IPv6List        []string        `json:"ipv6List"`
	LastSeen        float64         `json:"lastSeen"`
	PortList        []GatewayPort   `json:"portConfigs"`

	// WebAPI fields
	HardwareVersion string
}

func (g Gateway) GetType() string {
	return g.DeviceType.String()
}
func (g Gateway) GetMacAddress() string {
	return g.MacAddress
}
func (g Gateway) GetName() string {
	return g.Name
}
func (g Gateway) GetIP() string {
	return g.IP
}
func (g Gateway) GetModel() string {
	return g.Model
}
func (g Gateway) GetHardwareVersion() string {
	return g.HardwareVersion
}
func (g Gateway) GetFirmware() string {
	return g.FirmwareVersion
}
func (g Gateway) GetCpuUsage() float64 {
	return float64(g.CpuUsage)
}
func (g Gateway) GetMemUsage() float64 {
	return float64(g.RamUsage)
}
func (g Gateway) GetTemperature() float64 {
	return float64(g.Temperature)
}
func (g Gateway) GetLastSeen() float64 {
	return g.LastSeen
}
func (g Gateway) GetPorts() []Interface.Port {
	return Interface.ConvertToPortInterface(g.PortList)
}
func (g Gateway) GetRadios() []Interface.Radio {
	// Gateways do not have radios
	return nil
}
