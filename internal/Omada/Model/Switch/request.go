package Switch

import (
	"omada_exporter_go/internal/Omada/HttpClient/ApiClient"
	"omada_exporter_go/internal/Omada/Model/Devices"
)

func Get(devices []Devices.Device) (*[]Switch, error) {
	client := ApiClient.GetApiClient()

	var allData []Switch

	for _, d := range devices {
		if d.Type != Devices.DeviceType_Switch {
			continue
		}
		result, err := ApiClient.Get[Switch](*client, PATH_SWITCH, map[string]string{"switchMac": d.MacAddress}, nil, false)
		if err != nil {
			return nil, err
		}

		// Set the device type and name for each switch, based on device list
		for i := range *result {
			(*result)[i].DeviceType = Devices.DeviceType_Switch
			(*result)[i].Name = d.Name
		}

		allData = append(allData, *result...)
	}

	return &allData, nil
}
