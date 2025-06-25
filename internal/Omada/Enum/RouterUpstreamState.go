package Enum

import "fmt"

type RouterUpstreamState int8

const (
	RouterUpstreamState_PortDisabled RouterUpstreamState = -2
	RouterUpstreamState_LAN_Port     RouterUpstreamState = -1
	RouterUpstreamState_No           RouterUpstreamState = 0
	RouterUpstreamState_Yes          RouterUpstreamState = 1
)

func (od RouterUpstreamState) String() string {
	switch od {
	case RouterUpstreamState_PortDisabled:
		return "PortDisabled"
	case RouterUpstreamState_LAN_Port:
		return "LAN_Port"
	case RouterUpstreamState_No:
		return "No"
	case RouterUpstreamState_Yes:
		return "Yes"
	default:
		return "invalid"
	}
}
func (od RouterUpstreamState) Int() int64 {
	return int64(od)
}

func GetRouterUpstreamStatePossibleValues() string {
	return fmt.Sprintf(
		"%d - %s, %d - %s, %d - %s, %d - %s",
		RouterUpstreamState_PortDisabled.Int(),
		RouterUpstreamState_PortDisabled.String(),
		RouterUpstreamState_LAN_Port.Int(),
		RouterUpstreamState_LAN_Port.String(),
		RouterUpstreamState_No.Int(),
		RouterUpstreamState_No.String(),
		RouterUpstreamState_Yes.Int(),
		RouterUpstreamState_Yes.String(),
	)
}
