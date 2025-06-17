package main

import (
	"fmt"

	"omada_exporter_go/internal/Omada/Model/Devices"
	"omada_exporter_go/internal/Omada/Model/Gateway"
)

func main() {
	devices, err := Devices.Get()
	if err != nil {
		fmt.Println("Error fetching devices:", err)
		return
	}

	// switches, err := Switch.Get(*devices)
	// if err != nil {
	// 	fmt.Println("Error fetching switches:", err)
	// 	return
	// }

	// aps, err := AccessPoint.Get(*devices)
	// if err != nil {
	// 	fmt.Println("Error fetching access points:", err)
	// 	return
	// }

	router, err := Gateway.Get(*devices)
	if err != nil {
		fmt.Println("Error fetching gateways:", err)
		return
	}

	for _, r := range *router {
		fmt.Printf("Gateway: %+v\n", r.PortList[0].TransmitBytes)
	}

	// for _, ap := range *aps {
	// 	fmt.Printf("Access Point: %+v\n", ap)
	// }
	// for _, s := range *switches {
	// 	fmt.Printf("Switch: %+v\n", s)
	// }
}
