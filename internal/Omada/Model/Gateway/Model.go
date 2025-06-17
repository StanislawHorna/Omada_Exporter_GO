package Gateway

import (
	"fmt"

	"omada_exporter_go/internal/Omada/Enum"
)

const path_OpenApiGateway = "/openapi/v1/{omadaID}/sites/{siteID}/gateways/{gatewayMac}"
const path_WebApiGatewayPort = "{omadaID}/api/v2/sites/{siteID}/gateways/{gatewayMac}"

type rawGatewayPort struct {
	Port          int                       `json:"port"`
	PortName      string                    `json:"name"`
	PortDesc      string                    `json:"portDesc"`
	Mode          Enum.GatewayPortMode      `json:"mode"`
	IP            string                    `json:"ip"`
	Poe           bool                      `json:"poe"`
	LinkStatus    Enum.LinkStatus           `json:"status"`
	InternetState Enum.GatewayInternetState `json:"internetState"`
	Online        Enum.OnlineDetection      `json:"onlineDetection"`
	LinkSpeed     Enum.LinkSpeed            `json:"speed"`
	Duplex        Enum.DuplexMode           `json:"duplex"`
	Tx            int64                     `json:"tx"`
	Rx            int64                     `json:"rx"`
	Protocol      string                    `json:"proto"`
	// wanPortIpv6Config `json:"wanPortIpv6Config"`
	// wanPortIpv4Config `json:"wanPortIpv4Config"`
	Latency int     `json:"latency"`
	Loss    float32 `json:"loss"`
}

type rawGateway struct {
	PortStats []rawGatewayPort `json:"portStats"`
}

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
	Latency         int
	Loss            float32
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
	gp.ReceiveBytes = toMerge.Rx
	gp.TransmitBytes = toMerge.Tx
	gp.Latency = toMerge.Latency
	gp.Loss = toMerge.Loss

	return nil
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
	LastSeen        int64           `json:"lastSeen"`
	PortList        []GatewayPort   `json:"portConfigs"`
}
