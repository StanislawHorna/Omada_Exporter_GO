package AccessPoint

import (
	"omada_exporter_go/internal/Log"
	"omada_exporter_go/internal/Omada/Enum"
	"omada_exporter_go/internal/Omada/HttpClient/ApiClient"
	"omada_exporter_go/internal/Omada/HttpClient/WebClient"
	"omada_exporter_go/internal/Omada/Model/Devices"
)

func Get(devices []Devices.Device) (*[]AccessPoint, error) {
	Log.Debug("Fetching access points data")
	client := ApiClient.GetInstance()

	var allData []AccessPoint

	for _, d := range devices {
		if d.Type != Enum.DeviceType_AccessPoint {
			continue
		}
		result, err := ApiClient.Get[AccessPoint](client, path_OpenApiAccessPoint, map[string]string{"apMac": d.MacAddress}, nil, false)
		// If OpenAPI data is not available,
		// return error because OpenAPI data is base information, completed with WebAPI data
		if err != nil {
			return nil, Log.Error(err, "Failed to get access point data for AP %s", d.MacAddress)
		}

		// Set the device type and name for each access point, based on device list
		for i := range *result {
			(*result)[i].DeviceType = Enum.DeviceType_AccessPoint
			(*result)[i].Name = d.Name
			(*result)[i].LastSeen = d.LastSeen

			webApiData, err := getWebApiData(d)
			if err != nil {
				Log.Error(err, "Failed to get web API data for AP %s", d.MacAddress)
			} else {
				(*result)[i].merge(webApiData)
			}

			if err := getOpenApiRadioData(&(*result)[i]); err != nil {
				Log.Error(err, "Failed to get radio data for AP %s", d.MacAddress)
			}
		}

		allData = append(allData, *result...)
	}

	Log.Debug("Fetched %d access points", len(allData))

	return &allData, nil
}

func getWebApiData(d Devices.Device) (*webApiAccessPoint, error) {
	client := WebClient.GetInstance()

	result, err := WebClient.GetObject[webApiAccessPoint](client, path_WebApiAccessPointPort, map[string]string{"apMac": d.MacAddress}, nil)
	if err != nil {
		return nil, err
	}

	return (result), nil
}

func getOpenApiRadioData(ap *AccessPoint) error {
	client := ApiClient.GetInstance()

	result, err := ApiClient.Get[rawAccessPointRadio](client, path_OpenApiAccessPointRadio, map[string]string{"apMac": ap.MacAddress}, nil, false)
	if err != nil {
		return Log.Error(err, "Failed to get access point radio data for AP %s", ap.MacAddress)
	}
	// ApiClient.Get returns a slice of rawAccessPointRadio, but we expect a single item
	radioData := (*result)[0]
	ap.RadioList = radioData.ConvertToAccessPointRadio()

	return nil
}
