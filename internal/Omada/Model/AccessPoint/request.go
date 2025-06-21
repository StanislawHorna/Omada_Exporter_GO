package AccessPoint

import (
	"omada_exporter_go/internal/Omada/Enum"
	"omada_exporter_go/internal/Omada/HttpClient/ApiClient"
	"omada_exporter_go/internal/Omada/Model/Devices"
)

func Get(devices []Devices.Device) (*[]AccessPoint, error) {
	client := ApiClient.GetInstance()

	var allData []AccessPoint

	for _, d := range devices {
		if d.Type != Enum.DeviceType_AccessPoint {
			continue
		}
		result, err := ApiClient.Get[AccessPoint](*client, path_OpenApiAccessPoint, map[string]string{"apMac": d.MacAddress}, nil, false)
		if err != nil {
			return nil, err
		}

		// Set the device type and name for each access point, based on device list
		for i := range *result {
			(*result)[i].DeviceType = Enum.DeviceType_AccessPoint
			(*result)[i].Name = d.Name
			(*result)[i].LastSeen = d.LastSeen
		}

		allData = append(allData, *result...)
	}

	return &allData, nil
}
