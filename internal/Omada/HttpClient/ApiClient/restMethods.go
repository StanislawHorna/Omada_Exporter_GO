package ApiClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	Utils "omada_exporter_go/internal/Omada/HttpClient/Utils"
)

func Get[T any](client ApiClient, endpoint string, endpointPlaceholders map[string]string, queryParams map[string]string) (*[]T, error) {
	endpointPlaceholders = client.fillInOmadaIDs(endpointPlaceholders)
	endpoint = Utils.FillInEndpointPlaceholders(endpoint, endpointPlaceholders)
	if endpoint == "" {
		fmt.Println("Endpoint cannot be empty")
		return nil, fmt.Errorf("endpoint cannot be empty")
	}

	var allData []T
	currentPage := 1

	for {
		queryParamsWithPage := AddPaginationParams(queryParams, currentPage)
		url, err := Utils.CreateURL(client.BaseURL, endpoint, queryParamsWithPage)
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

		response, err := client.Http.Do(req)
		if err != nil {
			fmt.Println("Error making GET request:", err)
			return nil, err
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			fmt.Printf("Error: received status code %d from API\n", response.StatusCode)
			return nil, fmt.Errorf("non-OK status code: %d", response.StatusCode)
		}

		var apiResponse Response[Page[T]]
		if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}

		allData = append(allData, apiResponse.Result.Data...)

		if !apiResponse.Result.HasMorePages() {
			break
		}
		currentPage = apiResponse.Result.CurrentPage + 1
	}

	return &allData, nil
}
