package Enum

type DuplexMode int8

const (
	DuplexMode_Down DuplexMode = -1
	DuplexMode_Auto DuplexMode = 0
	DuplexMode_Half DuplexMode = 1
	DuplexMode_Full DuplexMode = 2
)

func (dm DuplexMode) String() string {
	switch dm {
	case DuplexMode_Down:
		return "Down"
	case DuplexMode_Auto:
		return "Auto"
	case DuplexMode_Half:
		return "Half"
	case DuplexMode_Full:
		return "Full"
	default:
		return "invalid"
	}
}
