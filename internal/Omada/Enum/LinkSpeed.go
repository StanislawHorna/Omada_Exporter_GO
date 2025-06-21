package Enum

type LinkSpeed int8

var (
	megabit_multiplier int64 = 1000 * 1000
	gigabit_multiplier int64 = 1000 * megabit_multiplier
)

const (
	LinkSpeed_Disabled LinkSpeed = -1
	LinkSpeed_Auto     LinkSpeed = 0
	LinkSpeed_10M      LinkSpeed = 1
	LinkSpeed_100M     LinkSpeed = 2
	LinkSpeed_1G       LinkSpeed = 3
	LinkSpeed_2_5G     LinkSpeed = 4
	LinkSpeed_10G      LinkSpeed = 5
	LinkSpeed_5G       LinkSpeed = 6
)

func (ls LinkSpeed) String() string {
	switch ls {
	case LinkSpeed_Disabled:
		return "Disabled"
	case LinkSpeed_Auto:
		return "Auto"
	case LinkSpeed_10M:
		return "10M"
	case LinkSpeed_100M:
		return "100M"
	case LinkSpeed_1G:
		return "1G"
	case LinkSpeed_2_5G:
		return "2.5G"
	case LinkSpeed_10G:
		return "10G"
	case LinkSpeed_5G:
		return "5G"
	default:
		return "invalid"
	}
}

func (ls LinkSpeed) Int() int64 {
	switch ls {
	case LinkSpeed_Disabled:
		return -1
	case LinkSpeed_Auto:
		return 0
	case LinkSpeed_10M:
		return 10 * megabit_multiplier
	case LinkSpeed_100M:
		return 100 * megabit_multiplier
	case LinkSpeed_1G:
		return 1000 * megabit_multiplier
	case LinkSpeed_2_5G:
		return 2500 * megabit_multiplier
	case LinkSpeed_10G:
		return 10 * gigabit_multiplier
	case LinkSpeed_5G:
		return 5 * gigabit_multiplier
	default:
		return -1
	}
}
