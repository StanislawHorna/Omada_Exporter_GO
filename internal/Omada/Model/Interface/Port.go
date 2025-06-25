package Interface

type Port interface {

	// Getters for port properties
	GetID() string
	GetPortName() string
	GetPortSpeed() float64
	GetPortDuplex() float64
	GetPortIP() string
	GetPortMode() string

	GetTxBytes() float64
	GetRxBytes() float64

	// Getters for specific port properties
	GetInternetState() float64
	GetInternetLatency() float64
	GetInternetLoss() float64
	GetUpstreamState() float64
	GetPortProtocol() string
}

func ConvertToPortInterface[T Port](portsToConvert []T) []Port {
	// The actual implementation would depend on the specific type of Port being used.
	ports := make([]Port, len(portsToConvert))
	for i, p := range portsToConvert {
		ports[i] = p // assign each specific port type to a Port interface
	}
	return ports
}
