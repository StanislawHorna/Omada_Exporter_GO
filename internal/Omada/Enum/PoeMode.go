package Enum

type PoeMode int8

const (
	PoeMode_off PoeMode = 0
	PoeMode_on  PoeMode = 1
)

func (pm PoeMode) String() string {
	switch pm {
	case PoeMode_off:
		return "Off"
	case PoeMode_on:
		return "On"
	default:
		return "invalid"
	}
}
