package Interface

type Radio interface {
	GetFrequency() string

	GetTxBytes() float64
	GetRxBytes() float64
}

func ConvertToRadioInterface[T Radio](radiosToConvert []T) []Radio {
	// The actual implementation would depend on the specific type of Port being used.
	radios := make([]Radio, len(radiosToConvert))
	for i, r := range radiosToConvert {
		radios[i] = r // assign each specific port type to a Port interface
	}
	return radios
}
