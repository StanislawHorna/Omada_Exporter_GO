package WebClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"omada_exporter_go/internal/Omada/HttpClient/Utils"
)

func GetObject[T any](client WebClient, endpoint string, endpointPlaceholders map[string]string, queryParams map[string]string) (*T, error) {
	endpointPlaceholders = client.fillInOmadaIDs(endpointPlaceholders)
	endpoint = Utils.FillInEndpointPlaceholders(endpoint, endpointPlaceholders)

	if endpoint == "" {
		fmt.Println("Endpoint cannot be empty")
		return nil, fmt.Errorf("endpoint cannot be empty")
	}

	queryParams = Utils.AddTimestampParam(queryParams)
	url, err := Utils.CreateURL(client.BaseURL, endpoint, queryParams)
	if err != nil {
		fmt.Println("Error creating URL:", err)
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	if err := client.setAuthorizationHeader(req); err != nil {
		fmt.Println("Error setting authorization header:", err)
		return nil, err
	}

	response, err := client.Client.Do(req)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d from API\n", response.StatusCode)
		return nil, fmt.Errorf("non-OK status code: %d", response.StatusCode)
	}

	var apiResponse Response[T]
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		fmt.Println("Error decoding response body:", err)
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	if apiResponse.ErrorCode != 0 {
		fmt.Printf("API error: %s (code %d)\n", apiResponse.Message, apiResponse.ErrorCode)
		return nil, fmt.Errorf("API error: %s (code %d)", apiResponse.Message, apiResponse.ErrorCode)
	}
	return &apiResponse.Result, nil
}
