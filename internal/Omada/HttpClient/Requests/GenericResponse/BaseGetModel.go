package GenericResponse

import (
	"encoding/json"
	"net/http"
	utils "omada_exporter_go/internal/Omada/HttpClient/Utils"
)

type OmadaGetResponse interface {
	Path(placeholders map[string]string) string
}

type OmadaResponse[T OmadaGetResponse] struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Result    T      `json:"result"`
}

func Get[T OmadaGetResponse](client *http.Client, baseUrl string, urlPlaceholders map[string]string, queryParams map[string]string) (*T, error) {
	var tempModel T
	url, err := utils.CreateURL(baseUrl, tempModel.Path(urlPlaceholders), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response OmadaResponse[T]
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response.Result, nil

}
