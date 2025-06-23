package Switch

import (
	"fmt"

	"omada_exporter_go/internal/Omada/Enum"
	"omada_exporter_go/internal/Omada/Model/Interface"
)

const path_OpenApiSwitch = "/openapi/v1/{omadaID}/sites/{siteID}/switches/{switchMac}"
const path_WebApiSwitchPort = "{omadaID}/api/v2/sites/{siteID}/switches/{switchMac}/ports"

type rawSwitchPortStatus struct {
	Port          int             `json:"port"`
	LinkStatus    Enum.LinkStatus `json:"linkStatus"`
	LinkSpeed     Enum.LinkSpeed  `json:"linkSpeed"`
	Duplex        Enum.DuplexMode `json:"duplex"`
	Poe           bool            `json:"poe"`
	Transmit      int64           `json:"tx"`
	Receive       int64           `json:"rx"`
	StpDiscarding bool            `json:"stpDiscarding"`
}

type rawSwitchPort struct {
	Port         int                 `json:"port"`
	ProfileName  string              `json:"profileName"`
	Disabled     bool                `json:"disabled"`
	MaxLinkSpeed Enum.LinkSpeed      `json:"maxSpeed"`
	PortStatus   rawSwitchPortStatus `json:"portStatus"`
}

// Implements Interface.Port
type SwitchPort struct {
	// OpenAPI fields
	Port                   int             `json:"port"`
	PortName               string          `json:"name"`
	ProfileID              string          `json:"profileId"`
	ProfileName            string          `json:"profileName"`
	ProfileOverrideEnabled bool            `json:"profileOverrideEnabled"`
	PoeMode                Enum.PoeMode    `json:"poeMode"`
	LagPort                bool            `json:"lagPort"`
	Status                 Enum.PortStatus `json:"status"`

	// WebAPI fields
	Disabled      bool
	LinkSpeed     Enum.LinkSpeed
	LinkStatus    Enum.LinkStatus
	MaxLinkSpeed  Enum.LinkSpeed
	DuplexMode    Enum.DuplexMode
	Poe           bool
	ReceiveBytes  int64
	TransmitBytes int64
}

func (sp *SwitchPort) merge(toMerge rawSwitchPort) error {
	if sp.Port != toMerge.Port {
		return fmt.Errorf("cannot merge SwitchPort with different port numbers: %d != %d", sp.Port, toMerge.Port)
	}
	sp.Disabled = toMerge.Disabled
	sp.LinkSpeed = toMerge.PortStatus.LinkSpeed
	sp.LinkStatus = toMerge.PortStatus.LinkStatus
	sp.MaxLinkSpeed = toMerge.MaxLinkSpeed
	sp.DuplexMode = toMerge.PortStatus.Duplex
	sp.Poe = toMerge.PortStatus.Poe
	sp.ReceiveBytes = toMerge.PortStatus.Receive
	sp.TransmitBytes = toMerge.PortStatus.Transmit

	return nil
}
func (sp SwitchPort) GetID() string {
	return fmt.Sprintf("%d", sp.Port)
}
func (sp SwitchPort) GetRxBytes() float64 {
	return float64(sp.ReceiveBytes)
}
func (sp SwitchPort) GetTxBytes() float64 {
	return float64(sp.TransmitBytes)
}

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
	return -1
}
func (s Switch) GetLastSeen() float64 {
	return s.LastSeen
}

// Implements Interface.Port
func (s Switch) GetPorts() []Interface.Port {
	return Interface.ConvertToPortInterface(s.PortList)
}
