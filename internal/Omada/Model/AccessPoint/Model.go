package AccessPoint

import (
	"omada_exporter_go/internal/Omada/Enum"
)

const path_OpenApiAccessPoint = "/openapi/v1/{omadaID}/sites/{siteID}/aps/{apMac}"

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
	DeviceType         Enum.DeviceType    `json:"deviceType"`
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
	LastSeen           float64
}

func (ap AccessPoint) GetType() string {
	return ap.DeviceType.String()
}
func (ap AccessPoint) GetMacAddress() string {
	return ap.MacAddress
}
func (ap AccessPoint) GetName() string {
	return ap.Name
}
func (ap AccessPoint) GetIP() string {
	return ap.IP
}
func (ap AccessPoint) GetModel() string {
	return ap.Model
}
func (ap AccessPoint) GetFirmware() string {
	return ap.FirmwareVersion
}
func (ap AccessPoint) GetCpuUsage() float64 {
	return float64(ap.CpuUsage)
}
func (ap AccessPoint) GetMemUsage() float64 {
	return float64(ap.RamUsage)
}
func (ap AccessPoint) GetTemperature() float64 {
	// Access Points do not provide temperature data
	return -1
}
func (ap AccessPoint) GetLastSeen() float64 {
	return ap.LastSeen
}
