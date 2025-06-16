package Devices

import "omada_exporter_go/internal/Omada/HttpClient/ApiClient"

func Get() (*[]Device, error) {
	client := ApiClient.GetApiClient()

	result, err := ApiClient.Get[Device](*client, PATH_DEVICES_LIST, nil, nil, true)
	if err != nil {
		return nil, err
	}
	return result, nil
}
