package Enum

type OnlineDetection int8

const (
	OnlineDetection_PortDisabled OnlineDetection = -2
	OnlineDetection_LAN_Port     OnlineDetection = -1
	OnlineDetection_No           OnlineDetection = 0
	OnlineDetection_Yes          OnlineDetection = 1
)

func (od OnlineDetection) String() string {
	switch od {
	case OnlineDetection_PortDisabled:
		return "PortDisabled"
	case OnlineDetection_LAN_Port:
		return "LAN_Port"
	case OnlineDetection_No:
		return "No"
	case OnlineDetection_Yes:
		return "Yes"
	default:
		return "invalid"
	}
}
