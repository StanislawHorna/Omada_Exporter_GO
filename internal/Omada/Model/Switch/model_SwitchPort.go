package Switch

import (
	"fmt"

	"omada_exporter_go/internal/Omada/Enum"
)

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

func (sp SwitchPort) GetID() string {
	return fmt.Sprintf("%d", sp.Port)
}
func (sp SwitchPort) GetRxBytes() float64 {
	return float64(sp.ReceiveBytes)
}
func (sp SwitchPort) GetTxBytes() float64 {
	return float64(sp.TransmitBytes)
}
func (sp SwitchPort) GetPortName() string {
	return sp.PortName
}
func (sp SwitchPort) GetPortSpeed() float64 {
	return float64(sp.LinkSpeed.Int())
}
func (sp SwitchPort) GetPortDuplex() float64 {
	return float64(sp.DuplexMode.Int())
}
func (sp SwitchPort) GetPortIP() string {
	// Switch ports do not have a specific IP address, return empty string
	return Enum.NotApplicable_String
}
func (sp SwitchPort) GetPortProtocol() string {
	// Access Point ports do not have a specific protocol, return empty string
	return Enum.NotApplicable_String
}
func (sp SwitchPort) GetPortMode() string {
	// Switch ports do not have a specific mode, return empty string
	return Enum.NotApplicable_String
}
func (sp SwitchPort) GetInternetState() float64 {
	// Switch ports do not have an internet state, return NotApplicable
	return Enum.NotApplicable_Float
}
func (sp SwitchPort) GetUpstreamState() float64 {
	// Switch ports do not have an upstream state, return NotApplicable
	return Enum.NotApplicable_Float
}
func (sp SwitchPort) GetInternetLatency() float64 {
	return Enum.NotApplicable_Float
}
func (sp SwitchPort) GetInternetLoss() float64 {
	return Enum.NotApplicable_Float
}
func (sp *SwitchPort) merge(toMerge webApiSwitchPort) error {
	if sp.Port != toMerge.Port {
		return fmt.Errorf("cannot merge SwitchPort with different port numbers: %d != %d", sp.Port, toMerge.Port)
	}
	sp.Disabled = toMerge.Disabled
	sp.LinkSpeed = toMerge.PortStatus.LinkSpeed
	sp.LinkStatus = toMerge.PortStatus.LinkStatus
	sp.MaxLinkSpeed = toMerge.MaxLinkSpeed
	sp.DuplexMode = toMerge.PortStatus.Duplex
	sp.Poe = toMerge.PortStatus.Poe

	// Omada webAPI for some reason returns port tx and rx in bits not bytes
	sp.ReceiveBytes = toMerge.PortStatus.Receive
	sp.TransmitBytes = toMerge.PortStatus.Transmit

	return nil
}
