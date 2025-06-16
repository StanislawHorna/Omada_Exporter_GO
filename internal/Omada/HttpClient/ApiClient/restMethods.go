package ApiClient

import (
	"encoding/json"
	"fmt"
	"net/http"

	Utils "omada_exporter_go/internal/Omada/HttpClient/Utils"
)

func Get[T any](client ApiClient, endpoint string, endpointPlaceholders map[string]string, queryParams map[string]string, usePagination bool) (*[]T, error) {
	endpointPlaceholders = client.fillInOmadaIDs(endpointPlaceholders)
	endpoint = Utils.FillInEndpointPlaceholders(endpoint, endpointPlaceholders)
	if endpoint == "" {
		fmt.Println("Endpoint cannot be empty")
		return nil, fmt.Errorf("endpoint cannot be empty")
	}

	var allData []T
	currentPage := 1

	for {
		var queryParamsToEncode map[string]string
		if usePagination {
			queryParamsToEncode = AddPaginationParams(queryParams, currentPage)
		} else {
			queryParamsToEncode = queryParams
		}

		url, err := Utils.CreateURL(client.BaseURL, endpoint, queryParamsToEncode)
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

		if response.StatusCode != http.StatusOK {
			fmt.Printf("Error: received status code %d from API\n", response.StatusCode)
			return nil, fmt.Errorf("non-OK status code: %d", response.StatusCode)
		}

		defer response.Body.Close()
		var nextPage int
		var data *[]T
		if usePagination {
			data, nextPage, err = decodePagedBody[T](response)
			if err != nil {
				fmt.Println("Error decoding response body:", err)
				return nil, err
			}
		} else {
			data, nextPage, err = decodeLongBody[T](response)
			if err != nil {
				fmt.Println("Error decoding response body:", err)
				return nil, err
			}
		}
		allData = append(allData, *data...)

		if nextPage <= 0 || !usePagination {
			break
		}
		currentPage = nextPage
	}

	return &allData, nil
}

func decodePagedBody[T any](response *http.Response) (*[]T, int, error) {
	var apiResponse Response[Page[T]]
	nextPage := -1

	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return nil, nextPage, fmt.Errorf("failed to decode response: %w", err)
	}
	if apiResponse.ErrorCode != 0 {
		return nil, nextPage, fmt.Errorf("API error: %s (code %d)", apiResponse.Message, apiResponse.ErrorCode)
	}
	if apiResponse.Result.HasMorePages() {
		nextPage = apiResponse.Result.CurrentPage + 1
	}
	return &apiResponse.Result.Data, nextPage, nil
}

func decodeLongBody[T any](response *http.Response) (*[]T, int, error) {
	var apiResponse Response[T]
	nextPage := -1

	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return nil, nextPage, fmt.Errorf("failed to decode response: %w", err)
	}
	if apiResponse.ErrorCode != 0 {
		return nil, nextPage, fmt.Errorf("API error: %s (code %d)", apiResponse.Message, apiResponse.ErrorCode)
	}
	// Convert an single object into a slice of objects to align structure with paginated responses
	return &[]T{apiResponse.Result}, nextPage, nil
}
