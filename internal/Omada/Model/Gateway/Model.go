package Gateway

import (
	"fmt"

	"omada_exporter_go/internal/Omada/Enum"
	"omada_exporter_go/internal/Omada/Model/Interface"
)

const path_OpenApiGateway = "/openapi/v1/{omadaID}/sites/{siteID}/gateways/{gatewayMac}"
const path_WebApiGatewayPort = "{omadaID}/api/v2/sites/{siteID}/gateways/{gatewayMac}"

type rawWanPortIpv4Config struct {
	IP            string `json:"ip"`
	Gateway       string `json:"gateway"`
	Gateway2      string `json:"gateway2"`
	PrimaryDNS    string `json:"priDns"`
	SecondaryDNS  string `json:"sndDns"`
	PrimaryDNS2   string `json:"priDns2"`
	SecondaryDNS2 string `json:"sndDns2"`
}

type rawWanPortIpv6Config struct {
	Enable        int                       `json:"enable"`
	Address       string                    `json:"addr"`
	Gateway       string                    `json:"gateway"`
	PrimaryDNS    string                    `json:"priDns"`
	SecondaryDNS  string                    `json:"sndDns"`
	InternetState Enum.GatewayInternetState `json:"internetState"`
}

type rawGatewayPort struct {
	Port              int                       `json:"port"`
	PortName          string                    `json:"name"`
	PortDesc          string                    `json:"portDesc"`
	Mode              Enum.GatewayPortMode      `json:"mode"`
	IP                string                    `json:"ip"`
	Poe               bool                      `json:"poe"`
	LinkStatus        Enum.LinkStatus           `json:"status"`
	InternetState     Enum.GatewayInternetState `json:"internetState"`
	Online            Enum.OnlineDetection      `json:"onlineDetection"`
	LinkSpeed         Enum.LinkSpeed            `json:"speed"`
	Duplex            Enum.DuplexMode           `json:"duplex"`
	TxBytes           int64                     `json:"tx"`
	RxBytes           int64                     `json:"rx"`
	TxPackets         int64                     `json:"txPkt"`
	RxPackets         int64                     `json:"rxPkt"`
	Protocol          string                    `json:"proto"`
	WanPortIpv4Config rawWanPortIpv4Config      `json:"wanPortIpv4Config"`
	WanPortIpv6Config rawWanPortIpv6Config      `json:"wanPortIpv6Config"`
	Latency           int                       `json:"latency"`
	Loss              float32                   `json:"loss"`
}

type rawGateway struct {
	PortStats []rawGatewayPort `json:"portStats"`
}

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
	IPv4Config      rawWanPortIpv4Config
	IPv6Config      rawWanPortIpv6Config
}

func (gp *GatewayPort) merge(toMerge rawGatewayPort) error {
	if gp.Port != toMerge.Port {
		return fmt.Errorf("cannot merge GatewayPort with different port numbers: %d != %d", gp.Port, toMerge.Port)
	}
	gp.PortName = toMerge.PortName
	gp.PortDescription = toMerge.PortDesc
	gp.Mode = toMerge.Mode
	gp.IP = toMerge.IP
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
func (gp GatewayPort) GetID() string {
	return fmt.Sprintf("%d", gp.Port)
}
func (gp GatewayPort) GetRxBytes() float64 {
	return float64(gp.ReceiveBytes)
}
func (gp GatewayPort) GetTxBytes() float64 {
	return float64(gp.TransmitBytes)
}

type Gateway struct {
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
