package AccessPoint

import (
	"omada_exporter_go/internal/Omada/Enum"
	"omada_exporter_go/internal/Omada/Model/Interface"
)

const (
	path_OpenApiAccessPoint = "/openapi/v1/{omadaID}/sites/{siteID}/aps/{apMac}"
)

// Implements Interface.Device
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

	RadioList []AccessPointRadio

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
func (ap AccessPoint) GetHardwareVersion() string {
	return ap.HardwareVersion
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
	return Enum.NotApplicable_Float
}
func (ap AccessPoint) GetLastSeen() float64 {
	return ap.LastSeen
}
func (ap AccessPoint) GetPorts() []Interface.Port {
	return Interface.ConvertToPortInterface(ap.PortList)
}
func (ap AccessPoint) GetRadios() []Interface.Radio {
	return Interface.ConvertToRadioInterface(ap.RadioList)
}

func (ap *AccessPoint) merge(toMerge *webApiAccessPoint) {
	ap.HardwareVersion = toMerge.HardwareVersion
	ap.PortList = make([]AccessPointPort, 1)
	ap.PortList[0] = AccessPointPort{
		PortReceiveBytes:    toMerge.WiredUpLink.RxBytes,
		PortTransmitBytes:   toMerge.WiredUpLink.TxBytes,
		PortReceivePackets:  toMerge.WiredUpLink.RxPackets,
		PortTransmitPackets: toMerge.WiredUpLink.TxPackets,
	}

}
