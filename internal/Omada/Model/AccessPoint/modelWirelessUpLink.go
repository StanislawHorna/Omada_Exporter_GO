package AccessPoint

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
