package Enum

type DeviceType string

const (
	DeviceType_Switch      DeviceType = "switch"
	DeviceType_AccessPoint DeviceType = "ap"
	DeviceType_Gateway     DeviceType = "gateway"
)

func (dt DeviceType) String() string {
	switch dt {
	case DeviceType_Switch:
		return "Switch"
	case DeviceType_AccessPoint:
		return "AccessPoint"
	case DeviceType_Gateway:
		return "Gateway"
	default:
		return "unknown"
	}
}
