package Enum

import "fmt"

type UpgradeNeeded bool

const (
	UpgradeNeeded_true  UpgradeNeeded = true
	UpgradeNeeded_false UpgradeNeeded = false
)

var toFloatMap = map[UpgradeNeeded]float64{true: 1.0, false: 0.0}
var toIntMap = map[UpgradeNeeded]int{true: 1, false: 0}
var toStringMap = map[UpgradeNeeded]string{true: "true", false: "false"}

func (un UpgradeNeeded) Float() float64 {
	return toFloatMap[un]
}

func (un UpgradeNeeded) Int() int {
	return toIntMap[un]
}

func (un UpgradeNeeded) String() string {
	return toStringMap[un]
}

func GetUpgradeNeededPossibleValues() string {
	return fmt.Sprintf(
		"%d - %s, %d - %s",
		UpgradeNeeded_true.Int(),
		UpgradeNeeded_true.String(),
		UpgradeNeeded_false.Int(),
		UpgradeNeeded_false.String(),
	)
}
