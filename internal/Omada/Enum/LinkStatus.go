package Enum

type LinkStatus int8

const (
	LinkStatus_Unknown LinkStatus = -1
	LinkStatus_Down    LinkStatus = 0
	LinkStatus_Up      LinkStatus = 1
)

func (ls LinkStatus) String() string {
	switch ls {
	case LinkStatus_Unknown:
		return "unknown"
	case LinkStatus_Down:
		return "down"
	case LinkStatus_Up:
		return "up"
	default:
		return "invalid"
	}
}
