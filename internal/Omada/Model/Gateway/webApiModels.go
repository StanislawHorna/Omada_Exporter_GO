package Gateway

import "omada_exporter_go/internal/Omada/Enum"

const path_WebApiGatewayPort = "{omadaID}/api/v2/sites/{siteID}/gateways/{gatewayMac}"

type webApiWanPortIpv4Config struct {
	IP            string `json:"ip"`
	Gateway       string `json:"gateway"`
	Gateway2      string `json:"gateway2"`
	PrimaryDNS    string `json:"priDns"`
	SecondaryDNS  string `json:"sndDns"`
	PrimaryDNS2   string `json:"priDns2"`
	SecondaryDNS2 string `json:"sndDns2"`
}

type webApiWanPortIpv6Config struct {
	Enable        int                       `json:"enable"`
	Address       string                    `json:"addr"`
	Gateway       string                    `json:"gateway"`
	PrimaryDNS    string                    `json:"priDns"`
	SecondaryDNS  string                    `json:"sndDns"`
	InternetState Enum.GatewayInternetState `json:"internetState"`
}

type webApiGatewayPort struct {
	Port              int                       `json:"port"`
	PortName          string                    `json:"name"`
	PortDesc          string                    `json:"portDesc"`
	Mode              Enum.GatewayPortMode      `json:"mode"`
	IP                string                    `json:"ip"`
	Poe               bool                      `json:"poe"`
	LinkStatus        Enum.LinkStatus           `json:"status"`
	InternetState     Enum.GatewayInternetState `json:"internetState"`
	Online            Enum.RouterUpstreamState  `json:"onlineDetection"`
	LinkSpeed         Enum.LinkSpeed            `json:"speed"`
	Duplex            Enum.DuplexMode           `json:"duplex"`
	TxBytes           int64                     `json:"tx"`
	RxBytes           int64                     `json:"rx"`
	TxPackets         int64                     `json:"txPkt"`
	RxPackets         int64                     `json:"rxPkt"`
	Protocol          string                    `json:"proto"`
	WanPortIpv4Config webApiWanPortIpv4Config   `json:"wanPortIpv4Config"`
	WanPortIpv6Config webApiWanPortIpv6Config   `json:"wanPortIpv6Config"`
	Latency           int                       `json:"latency"`
	Loss              float32                   `json:"loss"`
}

type rawGateway struct {
	HardwareVersion string              `json:"hwVersion"`
	PortStats       []webApiGatewayPort `json:"portStats"`
}
