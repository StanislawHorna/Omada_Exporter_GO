package Enum

import "fmt"

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

func (gs GatewayInternetState) Int() int64 {
	switch gs {
	case GatewayInternetState_Offline:
		return 0
	case GatewayInternetState_Online:
		return 1
	default:
		return NotApplicable_Int
	}
}
func GetInternetStatePossibleValues() string {
	return fmt.Sprintf(
		"%d - %s, %d - %s",
		GatewayInternetState_Offline.Int(),
		GatewayInternetState_Offline.String(),
		GatewayInternetState_Online.Int(),
		GatewayInternetState_Online.String(),
	)
}
