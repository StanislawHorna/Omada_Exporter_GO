package Gateway

import "omada_exporter_go/internal/Omada/Model/Devices"

const PATH_GATEWAY = "/openapi/v1/{omadaID}/sites/{siteID}/gateways/{gatewayMac}"

type PortSpeed int8

const (
	PortSpeed_Auto PortSpeed = 0
	PortSpeed_10M  PortSpeed = 1
	PortSpeed_100M PortSpeed = 2
	PortSpeed_1G   PortSpeed = 3
	PortSpeed_2_5G PortSpeed = 4
	PortSpeed_10G  PortSpeed = 5
	PortSpeed_5G   PortSpeed = 6
)

type DuplexMode int8

const (
	DuplexMode_Auto DuplexMode = 0
	DuplexMode_Half DuplexMode = 1
	DuplexMode_Full DuplexMode = 2
)

type MirrorMode int8

const (
	MirrorMode_Ingress       MirrorMode = 0
	MirrorMode_Egress        MirrorMode = 1
	MirrorMode_IngressEgress MirrorMode = 2
)

type GatewayPort struct {
	Port          int        `json:"port"`
	LinkSpeed     PortSpeed  `json:"linkSpeed"`
	DuplexMode    DuplexMode `json:"duplexMode"`
	MirrorEnabled bool       `json:"mirrorEnabled"`
	MirroredPorts []int      `json:"mirroredPorts"`
	MirrorMode    MirrorMode `json:"mirrorMode"`
	PortMode      int8       `json:"pvid"`
}

type Gateway struct {
	DeviceType      Devices.DeviceType `json:"deviceType"`
	Name            string             `json:"name"`
	MacAddress      string             `json:"mac"`
	Model           string             `json:"showModel"`
	FirmwareVersion string             `json:"firmwareVersion"`
	IP              string             `json:"ip"`
	Uptime          string             `json:"uptime"`
	Temperature     int                `json:"temp"`
	CpuUsage        int                `json:"cpuUtil"`
	RamUsage        int                `json:"memUtil"`
	IPv6List        []string           `json:"ipv6List"`
	LastSeen        int64              `json:"lastSeen"`
	PortList        []GatewayPort      `json:"portConfigs"`
}
