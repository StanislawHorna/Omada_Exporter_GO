package AccessPoint

import (
	"omada_exporter_go/internal/Omada/Enum"
	"omada_exporter_go/internal/Omada/Model/Interface"
)

const path_OpenApiAccessPoint = "/openapi/v1/{omadaID}/sites/{siteID}/aps/{apMac}"
const path_WebApiAccessPointPort = "{omadaID}/api/v2/sites/{siteID}/eaps/{apMac}"

type rawLanPort struct {
	TxPackets int64 `json:"upPackets"`
	RxPackets int64 `json:"downPackets"`
	TxBytes   int64 `json:"upBytes"`
	RxBytes   int64 `json:"downBytes"`
}

type rawAccessPoint struct {
	HardwareVersion string     `json:"hwVersion"`
	WiredUpLink     rawLanPort `json:"wiredUplink"`
}

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

// Implements Interface.Port
type AccessPointPort struct {
	PortReceiveBytes    int64
	PortTransmitBytes   int64
	PortReceivePackets  int64
	PortTransmitPackets int64
}

func (app AccessPointPort) GetID() string {
	// Access Point Port does not have a specific port number, return 1
	return "1"
}
func (app AccessPointPort) GetRxBytes() float64 {
	return float64(app.PortReceiveBytes)
}
func (app AccessPointPort) GetTxBytes() float64 {
	return float64(app.PortTransmitBytes)
}

type AccessPoint struct {
	// OpenAPI fields
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

	// WebAPI fields
	PortList        []AccessPointPort
	HardwareVersion string
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
func (ap AccessPoint) GetPorts() []Interface.Port {
	return Interface.ConvertToPortInterface(ap.PortList)
}
