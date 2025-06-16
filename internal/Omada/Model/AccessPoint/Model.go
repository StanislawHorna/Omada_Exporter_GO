package AccessPoint

import (
	"omada_exporter_go/internal/Omada/Model/Devices"
)

const PATH_ACCESS_POINT = "/openapi/v1/{omadaID}/sites/{siteID}/aps/{apMac}"

type ApWirelessUpLink struct {
	UplinkMac   string `json:"uplinkMac"`
	Name        string `json:"name"`
	Channel     int    `json:"channel"`
	Rssi        int    `json:"rssi"`
	Snr         int    `json:"snr"`
	TxRate      string `json:"txRate"`
	RxRateInt   int    `json:"rxRateInt"`
	RxRate      string `json:"rxRate"`
	UpBytes     int    `json:"upBytes"`
	DownBytes   int    `json:"downBytes"`
	UpPackets   int    `json:"upPackets"`
	DownPackets int    `json:"downPackets"`
	Activity    int    `json:"activity"`
}

type AccessPoint struct {
	DeviceType         Devices.DeviceType `json:"deviceType"`
	Name               string             `json:"name"`
	MacAddress         string             `json:"mac"`
	IP                 string             `json:"ip"`
	IPv6List           []string           `json:"ipv6List"`
	Model              string             `json:"showModel"`
	WlanID             string             `json:"wlanId"`
	FirmwareVersion    string             `json:"firmwareVersion"`
	WirelessUplinkInfo []ApWirelessUpLink `json:"wirelessUplinkInfo"`
	CpuUsage           int                `json:"cpuUtil"`
	RamUsage           int                `json:"memUtil"`
	Uptime             int64              `json:"uptimeLong"`
}
