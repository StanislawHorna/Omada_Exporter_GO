package Enum

import "fmt"

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
		return NotApplicable_String
	}
}

func (dm DuplexMode) Int() int64 {
	switch dm {
	case DuplexMode_Down:
		return -1
	case DuplexMode_Auto:
		return 0
	case DuplexMode_Half:
		return 1
	case DuplexMode_Full:
		return 2
	default:
		return NotApplicable_Int
	}
}

func GetDuplexPossibleValues() string {
	return fmt.Sprintf(
		"%d - %s, %d - %s, %d - %s, %d - %s",
		DuplexMode_Down.Int(),
		DuplexMode_Down.String(),
		DuplexMode_Auto.Int(),
		DuplexMode_Auto.String(),
		DuplexMode_Half.Int(),
		DuplexMode_Half.String(),
		DuplexMode_Full.Int(),
		DuplexMode_Full.String(),
	)
}
