package Gateway

import (
	"fmt"

	"omada_exporter_go/internal/Omada/Enum"
	"omada_exporter_go/internal/Omada/HttpClient/ApiClient"
	"omada_exporter_go/internal/Omada/HttpClient/WebClient"
	"omada_exporter_go/internal/Omada/Model/Devices"
)

func Get(devices []Devices.Device) (*[]Gateway, error) {

	var allData []Gateway

	for _, d := range devices {
		if d.Type != Enum.DeviceType_Gateway {
			continue
		}
		openApiResult, err := getOpenApiData(d)
		if err != nil {
			return nil, err
		}

		webApiResult, err := getWebApiData(d)
		if err != nil {
			return nil, err
		}

		// reference to slice index 0, since we expect only one gateway per MAC address
		(*openApiResult)[0].HardwareVersion = webApiResult.HardwareVersion

		for i := range (*openApiResult)[0].PortList {
			// Merge the web API data into the OpenAPI result
			for _, webPort := range (*webApiResult).PortStats {
				if (*openApiResult)[0].PortList[i].Port == webPort.Port {
					if err := (*openApiResult)[0].PortList[i].merge(webPort); err != nil {
						fmt.Printf("Error merging port data for gateway %s: %v\n", d.MacAddress, err)
					}
					// If port is down set speed and duplex as disabled
					if (*openApiResult)[0].PortList[i].LinkStatus == Enum.LinkStatus_Down {
						(*openApiResult)[0].PortList[i].Mode = Enum.GatewayPortMode_Down
						(*openApiResult)[0].PortList[i].DuplexMode = Enum.DuplexMode_Down
						(*openApiResult)[0].PortList[i].Online = Enum.OnlineDetection_PortDisabled
						(*openApiResult)[0].PortList[i].LinkSpeed = Enum.LinkSpeed_Disabled
						(*openApiResult)[0].PortList[i].DuplexMode = Enum.DuplexMode_Down
						(*openApiResult)[0].PortList[i].Latency = 0
						(*openApiResult)[0].PortList[i].Loss = 1.0

					}
					if (*openApiResult)[0].PortList[i].Mode == Enum.GatewayPortMode_LAN {
						(*openApiResult)[0].PortList[i].Loss = 0.0  // Set loss to 0 for LAN ports
						(*openApiResult)[0].PortList[i].Latency = 0 // Set latency to 0 for LAN ports

						(*openApiResult)[0].PortList[i].Online = Enum.OnlineDetection_LAN_Port

					}
					break
				}
			}
		}

		allData = append(allData, *openApiResult...)
	}

	return &allData, nil
}

func getOpenApiData(d Devices.Device) (*[]Gateway, error) {
	client := ApiClient.GetInstance()

	result, err := ApiClient.Get[Gateway](*client, path_OpenApiGateway, map[string]string{"gatewayMac": d.MacAddress}, nil, false)
	if err != nil {
		return nil, err
	}

	if len(*result) == 0 {
		return nil, nil
	}

	// Set the device type and name based on device entry
	(*result)[0].DeviceType = Enum.DeviceType_Gateway
	(*result)[0].Name = d.Name

	return result, nil
}

func getWebApiData(d Devices.Device) (*rawGateway, error) {
	client := WebClient.GetInstance()

	result, err := WebClient.GetObject[rawGateway](*client, path_WebApiGatewayPort, map[string]string{"gatewayMac": d.MacAddress}, nil)
	if err != nil {
		return nil, err
	}

	if len((*result).PortStats) == 0 {
		return nil, nil
	}

	return result, nil
}
