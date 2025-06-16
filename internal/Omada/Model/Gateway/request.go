package Gateway

import (
	"omada_exporter_go/internal/Omada/HttpClient/ApiClient"
	"omada_exporter_go/internal/Omada/Model/Devices"
)

func Get(devices []Devices.Device) (*[]Gateway, error) {
	client := ApiClient.GetApiClient()

	var allData []Gateway

	for _, d := range devices {
		if d.Type != Devices.DeviceType_Router {
			continue
		}
		result, err := ApiClient.Get[Gateway](*client, PATH_GATEWAY, map[string]string{"gatewayMac": d.MacAddress}, nil, false)
		if err != nil {
			return nil, err
		}

		// Set the device type and name for each switch, based on device list
		for i := range *result {
			(*result)[i].DeviceType = Devices.DeviceType_Router
			(*result)[i].Name = d.Name
		}

		allData = append(allData, *result...)
	}

	return &allData, nil
}
