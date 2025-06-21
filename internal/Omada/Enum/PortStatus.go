package Enum

type PortStatus int8

const (
	PortStatus_disabled PortStatus = 0
	PortStatus_enabled  PortStatus = 1
)

func (ps PortStatus) String() string {
	switch ps {
	case PortStatus_disabled:
		return "Disabled"
	case PortStatus_enabled:
		return "Enabled"
	default:
		return "invalid"
	}
}
