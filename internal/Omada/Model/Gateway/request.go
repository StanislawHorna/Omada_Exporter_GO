package Gateway

import (
	"omada_exporter_go/internal/Log"
	"omada_exporter_go/internal/Omada/Enum"
	"omada_exporter_go/internal/Omada/HttpClient/ApiClient"
	"omada_exporter_go/internal/Omada/HttpClient/WebClient"
	"omada_exporter_go/internal/Omada/Model/Devices"
)

func Get(devices []Devices.Device) (*[]Gateway, error) {
	Log.Debug("Fetching gateways data")
	var allData []Gateway

	for _, d := range devices {
		if d.Type != Enum.DeviceType_Gateway {
			continue
		}
		openApiResult, err := getOpenApiData(d)
		if err != nil {
			return nil, Log.Error(err, "Failed to get OpenAPI data for gateway %s", d.MacAddress)
		}

		webApiResult, err := getWebApiData(d)
		if err != nil {
			return nil, Log.Error(err, "Failed to get WebAPI data for gateway %s", d.MacAddress)
		}

		// reference to slice index 0, since we expect only one gateway per MAC address
		(*openApiResult)[0].HardwareVersion = webApiResult.HardwareVersion

		for i := range (*openApiResult)[0].PortList {
			// Merge the web API data into the OpenAPI result
			for _, webPort := range (*webApiResult).PortStats {
				if (*openApiResult)[0].PortList[i].Port == webPort.Port {
					if err := (*openApiResult)[0].PortList[i].merge(webPort); err != nil {
						Log.Error(err, "Failed to merge port data for gateway %s", d.MacAddress)
					}
					// If port is down set speed and duplex as disabled
					if (*openApiResult)[0].PortList[i].LinkStatus == Enum.LinkStatus_Down {
						(*openApiResult)[0].PortList[i].Mode = Enum.GatewayPortMode_Down
						(*openApiResult)[0].PortList[i].DuplexMode = Enum.DuplexMode_Down
						(*openApiResult)[0].PortList[i].Online = Enum.RouterUpstreamState_PortDisabled
						(*openApiResult)[0].PortList[i].LinkSpeed = Enum.LinkSpeed_Disabled
						(*openApiResult)[0].PortList[i].DuplexMode = Enum.DuplexMode_Down
						(*openApiResult)[0].PortList[i].Latency = 0
						(*openApiResult)[0].PortList[i].Loss = 1.0

					}
					if (*openApiResult)[0].PortList[i].Mode == Enum.GatewayPortMode_LAN {
						(*openApiResult)[0].PortList[i].Loss = 0.0  // Set loss to 0 for LAN ports
						(*openApiResult)[0].PortList[i].Latency = 0 // Set latency to 0 for LAN ports
						(*openApiResult)[0].PortList[i].Online = Enum.RouterUpstreamState_LAN_Port

					}
					break
				}
			}
		}

		allData = append(allData, *openApiResult...)
	}

	Log.Info("Fetched %d gateways", len(allData))

	return &allData, nil
}

func getOpenApiData(d Devices.Device) (*[]Gateway, error) {
	client := ApiClient.GetInstance()

	result, err := ApiClient.Get[Gateway](*client, path_OpenApiGateway, map[string]string{"gatewayMac": d.MacAddress}, nil, false)
	if err != nil {
		return nil, err
	}

	if len(*result) == 0 {
		Log.Warn("No gateway data in OpenAPI found for device %s", d.MacAddress)
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
		Log.Warn("No gateway data in WebAPI found for device %s", d.MacAddress)
		return nil, nil
	}

	return result, nil
}
