package AccessPoint

import "omada_exporter_go/internal/Omada/Enum"

const path_OpenApiAccessPointRadio = "/openapi/v1/{omadaID}/sites/{siteID}/aps/{apMac}/radios"

type apRadioTraffic struct {
	Frequency   Enum.RadioFrequency
	RxPkts      float64 `json:"rxPkts"`
	TxPkts      float64 `json:"txPkts"`
	Rx          float64 `json:"rx"`
	Tx          float64 `json:"tx"`
	RxDropPkts  float64 `json:"rxDropPkts"`
	TxDropPkts  float64 `json:"txDropPkts"`
	RxErrPkts   float64 `json:"rxErrPkts"`
	TxErrPkts   float64 `json:"txErrPkts"`
	RxRetryPkts float64 `json:"rxRetryPkts"`
	TxRetryPkts float64 `json:"txRetryPkts"`
}

type apRadioConfig struct {
	Frequency     Enum.RadioFrequency
	ActualChannel string  `json:"actualChannel"`
	MaxTxRate     float64 `json:"maxTxRate"`
	Region        int     `json:"region"`
	Bandwidth     string  `json:"bandWidth"`
	Mode          string  `json:"rdMode"`
	TxUtil        float64 `json:"txUtil"`
	RxUtil        float64 `json:"rxUtil"`
	InterUtil     float64 `json:"interUtil"`
}

func mergeConfigAndTraffic(freq Enum.RadioFrequency, config apRadioConfig, traffic apRadioTraffic) AccessPointRadio {
	return AccessPointRadio{
		Frequency: freq,

		// Config
		ActualChannel: config.ActualChannel,
		MaxTxRate:     config.MaxTxRate,
		Region:        config.Region,
		Bandwidth:     config.Bandwidth,
		Mode:          config.Mode,
		TxUtil:        config.TxUtil,
		RxUtil:        config.RxUtil,
		InterUtil:     config.InterUtil,

		// Traffic
		ReceivePackets:  traffic.RxPkts,
		TransmitPackets: traffic.TxPkts,
		ReceiveBytes:    traffic.Rx,
		TransmitBytes:   traffic.Tx,
		RxDropPackets:   traffic.RxDropPkts,
		TxDropPackets:   traffic.TxDropPkts,
		RxErrPackets:    traffic.RxErrPkts,
		TxErrPackets:    traffic.TxErrPkts,
		RxRetryPackets:  traffic.RxRetryPkts,
		TxRetryPackets:  traffic.TxRetryPkts,
	}
}

type rawAccessPointRadio struct {
	Traffic24GHz apRadioTraffic `json:"radioTraffic2g"`
	Traffic50GHz apRadioTraffic `json:"radioTraffic5g"`
	Config24GHz  apRadioConfig  `json:"wp2g"`
	Config50GHz  apRadioConfig  `json:"wp5g"`
}

func (apr rawAccessPointRadio) ConvertToAccessPointRadio() []AccessPointRadio {
	apr.Traffic24GHz.Frequency = Enum.RadioFrequency_2_4_Ghz
	apr.Config24GHz.Frequency = Enum.RadioFrequency_2_4_Ghz
	apr.Traffic50GHz.Frequency = Enum.RadioFrequency_5_0_Ghz
	apr.Config50GHz.Frequency = Enum.RadioFrequency_5_0_Ghz

	return []AccessPointRadio{
		mergeConfigAndTraffic(Enum.RadioFrequency_2_4_Ghz, apr.Config24GHz, apr.Traffic24GHz),
		mergeConfigAndTraffic(Enum.RadioFrequency_5_0_Ghz, apr.Config50GHz, apr.Traffic50GHz),
	}
}

type AccessPointRadio struct {
	Frequency Enum.RadioFrequency

	// Configuration data
	ActualChannel string
	MaxTxRate     float64
	Region        int
	Bandwidth     string
	Mode          string
	TxUtil        float64
	RxUtil        float64
	InterUtil     float64

	// Traffic data
	ReceivePackets  float64
	TransmitPackets float64
	ReceiveBytes    float64
	TransmitBytes   float64
	RxDropPackets   float64
	TxDropPackets   float64
	RxErrPackets    float64
	TxErrPackets    float64
	RxRetryPackets  float64
	TxRetryPackets  float64
}

func (apr AccessPointRadio) GetFrequency() string {
	return apr.Frequency.String()
}
func (apr AccessPointRadio) GetTxBytes() float64 {
	return apr.TransmitBytes
}
func (apr AccessPointRadio) GetRxBytes() float64 {
	return apr.ReceiveBytes
}
