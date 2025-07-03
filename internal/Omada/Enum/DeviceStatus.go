package Enum

import "fmt"

type DeviceStatus int

const (
	DeviceStatus_disconnected     DeviceStatus = 0
	DeviceStatus_connected        DeviceStatus = 1
	DeviceStatus_pending          DeviceStatus = 2
	DeviceStatus_heartbeatMissing DeviceStatus = 3
	DeviceStatus_isolated         DeviceStatus = 4
)

func (ds DeviceStatus) String() string {
	switch ds {
	case DeviceStatus_disconnected:
		return "Disconnected"
	case DeviceStatus_connected:
		return "Connected"
	case DeviceStatus_pending:
		return "Pending"
	case DeviceStatus_heartbeatMissing:
		return "HeartbeatMissing"
	case DeviceStatus_isolated:
		return "Isolated"
	default:
		return "unknown"
	}
}
func (ds DeviceStatus) Int() int64 {
	return int64(ds)
}

func GetDeviceStatusPossibleValues() string {
	return fmt.Sprintf(
		"%d - %s, %d - %s, %d - %s, %d - %s, %d - %s",
		DeviceStatus_disconnected.Int(),
		DeviceStatus_disconnected.String(),
		DeviceStatus_connected.Int(),
		DeviceStatus_connected.String(),
		DeviceStatus_pending.Int(),
		DeviceStatus_pending.String(),
		DeviceStatus_heartbeatMissing.Int(),
		DeviceStatus_heartbeatMissing.String(),
		DeviceStatus_isolated.Int(),
		DeviceStatus_isolated.String(),
	)
}
