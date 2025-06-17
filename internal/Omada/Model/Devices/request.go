package Devices

import "omada_exporter_go/internal/Omada/HttpClient/ApiClient"

func Get() (*[]Device, error) {
	client := ApiClient.GetInstance()

	result, err := ApiClient.Get[Device](*client, path_OpenApiDevicesList, nil, nil, true)
	if err != nil {
		return nil, err
	}
	return result, nil
}
