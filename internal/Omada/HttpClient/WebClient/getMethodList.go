package WebClient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"omada_exporter_go/internal/Log"
	"omada_exporter_go/internal/Omada/HttpClient/Utils"
)

func GetList[T any](client *WebClient, endpoint string, endpointPlaceholders map[string]string, queryParams map[string]string, usePagination bool) (*[]T, error) {
	endpointPlaceholders = client.fillInOmadaIDs(endpointPlaceholders)
	endpoint = Utils.FillInEndpointPlaceholders(endpoint, endpointPlaceholders)

	if endpoint == "" {
		return nil, Log.Error(nil, "Endpoint is empty")
	}

	var allData []T
	currentPage := 1

	for {
		var queryParamsToEncode map[string]string
		if usePagination {
			queryParamsToEncode = AddPaginationParams(queryParams, currentPage)
			Log.Debug("Requesting page %d of endpoint %s with query params: %v", currentPage, endpoint, queryParamsToEncode)
		} else {
			queryParamsToEncode = queryParams
			Log.Debug("Requesting endpoint %s with query params: %v", endpoint, queryParamsToEncode)
		}

		queryParamsToEncode = Utils.AddTimestampParam(queryParamsToEncode)
		url, err := Utils.CreateURL(client.BaseURL, endpoint, queryParamsToEncode)
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
		if response.StatusCode != http.StatusOK {
			return nil, Log.Error(nil, "Received non-OK status code from API: %d", response.StatusCode)
		}

		defer response.Body.Close()
		var nextPage int
		var data *[]T
		if usePagination {
			data, nextPage, err = decodePagedBody[T](response)
			if err != nil {
				return nil, Log.Error(err, "")
			}
		} else {
			data, nextPage, err = decodeLongBody[T](response)
			if err != nil {
				return nil, Log.Error(err, "")
			}
		}
		allData = append(allData, *data...)

		if nextPage <= 0 || !usePagination {
			Log.Debug("No more pages to fetch or pagination not used, breaking the loop")
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
		return nil, nextPage, Log.Error(err, "Failed to decode paginated response")
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
	var apiResponse Response[[]T]
	nextPage := -1

	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		return nil, nextPage, Log.Error(err, "Failed to decode long response")
	}
	if apiResponse.ErrorCode != 0 {
		return nil, nextPage, Log.Error(nil, "API error: %s (code %d)", apiResponse.Message, apiResponse.ErrorCode)
	}
	// Convert an single object into a slice of objects to align structure with paginated responses
	return &apiResponse.Result, nextPage, nil
}
