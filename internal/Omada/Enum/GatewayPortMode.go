package Enum

type GatewayPortMode int8

const (
	GatewayPortMode_Down GatewayPortMode = -1
	GatewayPortMode_WAN  GatewayPortMode = 0
	GatewayPortMode_LAN  GatewayPortMode = 1
)

func (gpm GatewayPortMode) String() string {
	switch gpm {
	case GatewayPortMode_Down:
		return "Down"
	case GatewayPortMode_WAN:
		return "WAN"
	case GatewayPortMode_LAN:
		return "LAN"
	default:
		return "invalid"
	}
}
