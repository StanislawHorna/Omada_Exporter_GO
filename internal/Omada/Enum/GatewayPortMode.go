package Enum

type GatewayPortMode int8

const (
	GatewayPortMode_Unknown GatewayPortMode = -1
	GatewayPortMode_WAN     GatewayPortMode = 0
	GatewayPortMode_LAN     GatewayPortMode = 1
)
