package AccessPoint

import (
	"omada_exporter_go/internal/Omada/HttpClient/ApiClient"
	"omada_exporter_go/internal/Omada/Model/Devices"
)

func Get(devices []Devices.Device) (*[]AccessPoint, error) {
	client := ApiClient.GetApiClient()

	var allData []AccessPoint

	for _, d := range devices {
		if d.Type != Devices.DeviceType_AccessPoint {
			continue
		}
		result, err := ApiClient.Get[AccessPoint](*client, PATH_ACCESS_POINT, map[string]string{"apMac": d.MacAddress}, nil, false)
		if err != nil {
			return nil, err
		}

		// Set the device type and name for each access point, based on device list
		for i := range *result {
			(*result)[i].DeviceType = Devices.DeviceType_AccessPoint
			(*result)[i].Name = d.Name
		}

		allData = append(allData, *result...)
	}

	return &allData, nil
}
