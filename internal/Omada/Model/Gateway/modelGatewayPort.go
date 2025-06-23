package Gateway

import (
	"fmt"

	"omada_exporter_go/internal/Omada/Enum"
)

const path_OpenApiGateway = "/openapi/v1/{omadaID}/sites/{siteID}/gateways/{gatewayMac}"

// Implements Interface.Port
type GatewayPort struct {
	// OpenAPI fields
	Port          int             `json:"port"`
	LinkSpeed     Enum.LinkSpeed  `json:"linkSpeed"`
	DuplexMode    Enum.DuplexMode `json:"duplexMode"`
	MirrorEnabled bool            `json:"mirrorEnabled"`
	MirroredPorts []int           `json:"mirroredPorts"`
	MirrorMode    Enum.MirrorMode `json:"mirrorMode"`
	PortMode      int8            `json:"pvid"`

	// WebAPI fields
	PortName        string
	PortDescription string
	Mode            Enum.GatewayPortMode
	IP              string
	Protocol        string
	Poe             bool
	LinkStatus      Enum.LinkStatus
	InternetState   Enum.GatewayInternetState
	Online          Enum.OnlineDetection
	ReceiveBytes    int64
	TransmitBytes   int64
	ReceivePackets  int64
	TransmitPackets int64
	Latency         int
	Loss            float32
	IPv4Config      webApiWanPortIpv4Config
	IPv6Config      webApiWanPortIpv6Config
}

func (gp GatewayPort) GetID() string {
	return fmt.Sprintf("%d", gp.Port)
}
func (gp GatewayPort) GetRxBytes() float64 {
	return float64(gp.ReceiveBytes)
}
func (gp GatewayPort) GetTxBytes() float64 {
	return float64(gp.TransmitBytes)
}
func (gp GatewayPort) GetPortName() string {
	return gp.PortDescription
}
func (gp GatewayPort) GetPortSpeed() float64 {
	return float64(gp.LinkSpeed.Int())
}
func (gp GatewayPort) GetPortIP() string {
	return gp.IP
}
func (gp GatewayPort) GetPortProtocol() string {
	return gp.Protocol
}
func (gp *GatewayPort) merge(toMerge webApiGatewayPort) error {
	if gp.Port != toMerge.Port {
		return fmt.Errorf("cannot merge GatewayPort with different port numbers: %d != %d", gp.Port, toMerge.Port)
	}
	gp.PortName = toMerge.PortName
	gp.PortDescription = toMerge.PortDesc
	gp.Mode = toMerge.Mode
	gp.IP = toMerge.IP
	gp.Protocol = toMerge.Protocol
	gp.Poe = toMerge.Poe
	gp.LinkStatus = toMerge.LinkStatus
	gp.InternetState = toMerge.InternetState
	gp.Online = toMerge.Online
	gp.LinkSpeed = toMerge.LinkSpeed
	gp.DuplexMode = toMerge.Duplex
	gp.ReceiveBytes = toMerge.RxBytes
	gp.TransmitBytes = toMerge.TxBytes
	gp.ReceivePackets = toMerge.RxPackets
	gp.TransmitPackets = toMerge.TxPackets
	gp.Latency = toMerge.Latency
	gp.Loss = toMerge.Loss
	gp.IPv4Config = toMerge.WanPortIpv4Config
	gp.IPv6Config = toMerge.WanPortIpv6Config

	return nil
}
