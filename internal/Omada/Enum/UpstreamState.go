package Enum

import "fmt"

type UpstreamState int8

const (
	UpstreamState_PortDisabled UpstreamState = -2
	UpstreamState_LAN_Port     UpstreamState = -1
	UpstreamState_No           UpstreamState = 0
	UpstreamState_Yes          UpstreamState = 1
)

func (od UpstreamState) String() string {
	switch od {
	case UpstreamState_PortDisabled:
		return "PortDisabled"
	case UpstreamState_LAN_Port:
		return "LAN_Port"
	case UpstreamState_No:
		return "No"
	case UpstreamState_Yes:
		return "Yes"
	default:
		return "invalid"
	}
}
func (od UpstreamState) Int() int64 {
	return int64(od)
}

func GetUpstreamStatePossibleValues() string {
	return fmt.Sprintf(
		"%d - %s, %d - %s, %d - %s, %d - %s",
		UpstreamState_PortDisabled.Int(),
		UpstreamState_PortDisabled.String(),
		UpstreamState_LAN_Port.Int(),
		UpstreamState_LAN_Port.String(),
		UpstreamState_No.Int(),
		UpstreamState_No.String(),
		UpstreamState_Yes.Int(),
		UpstreamState_Yes.String(),
	)
}
