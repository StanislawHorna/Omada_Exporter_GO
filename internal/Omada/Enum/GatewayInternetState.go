package Enum

type GatewayInternetState int8

const (
	GatewayInternetState_Unknown GatewayInternetState = -1
	GatewayInternetState_Offline GatewayInternetState = 0
	GatewayInternetState_Online  GatewayInternetState = 1
)

func (gs GatewayInternetState) String() string {
	switch gs {
	case GatewayInternetState_Unknown:
		return "Unknown"
	case GatewayInternetState_Offline:
		return "Offline"
	case GatewayInternetState_Online:
		return "Online"
	default:
		return "invalid"
	}
}
