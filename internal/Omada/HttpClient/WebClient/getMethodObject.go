package WebClient

import (
	"encoding/json"
	"net/http"

	"omada_exporter_go/internal/Log"
	"omada_exporter_go/internal/Omada/HttpClient/Utils"
)

func GetObject[T any](client WebClient, endpoint string, endpointPlaceholders map[string]string, queryParams map[string]string) (*T, error) {
	endpointPlaceholders = client.fillInOmadaIDs(endpointPlaceholders)
	endpoint = Utils.FillInEndpointPlaceholders(endpoint, endpointPlaceholders)

	if endpoint == "" {
		return nil, Log.Error(nil, "Endpoint is empty")
	}

	queryParams = Utils.AddTimestampParam(queryParams)
	url, err := Utils.CreateURL(client.BaseURL, endpoint, queryParams)
	if err != nil {
		return nil, Log.Error(err, "Failed to create URL")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, Log.Error(err, "Failed to create HTTP request")
	}

	httpClient, err := client.getHttpClient()
	if err != nil {
		return nil, err
	}

	if err := client.setAuthorizationHeader(req); err != nil {
		return nil, Log.Error(err, "Failed to set authorization header")
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, Log.Error(err, "Failed to make GET request")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, Log.Error(nil, "Received non-OK status code from API: %d", response.StatusCode)
	}

	var apiResponse Response[T]
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return nil, Log.Error(err, "Failed to decode long response")
	}
	if apiResponse.ErrorCode != 0 {
		return nil, Log.Error(nil, "API error: %s (code %d)", apiResponse.Message, apiResponse.ErrorCode)
	}
	return &apiResponse.Result, nil
}
