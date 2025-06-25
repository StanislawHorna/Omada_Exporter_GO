package AccessPoint

const path_WebApiAccessPointPort = "{omadaID}/api/v2/sites/{siteID}/eaps/{apMac}"

type webApiLanPort struct {
	TxPackets int64 `json:"upPackets"`
	RxPackets int64 `json:"downPackets"`
	TxBytes   int64 `json:"upBytes"`
	RxBytes   int64 `json:"downBytes"`
}

type webApiAccessPoint struct {
	HardwareVersion string        `json:"hwVersion"`
	WiredUpLink     webApiLanPort `json:"wiredUplink"`
}
