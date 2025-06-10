package GenericResponse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	utils "omada_exporter_go/internal/Omada/HttpClient/Utils"
)

type OmadaPostResponseModel interface {
	Path(placeholders map[string]string) string
	Payload(data map[string]any) (any, error)
}

type OmadaPostResponse[T OmadaPostResponseModel] struct {
	ErrorCode int    `json:"errorCode"`
	ErrorMsg  string `json:"errorMsg"`
	Result    T      `json:"result"`
}

func Post[T OmadaPostResponseModel](
	client *http.Client,
	baseUrl string,
	urlPlaceholders map[string]string,
	queryParams map[string]string,
	payload map[string]string,
) (*T, error) {
	var tempModel T
	url, err := utils.CreateURL(baseUrl, tempModel.Path(urlPlaceholders), queryParams)
	if err != nil {
		return nil, err
	}
	// requestPayload, err := tempModel.Payload(payload)
	// if err != nil {
	// 	return nil, err
	// }
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	fmt.Println(payload)
	resp, err := client.Post(url, "application/json", bytes.NewReader(bodyBytes))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var response OmadaPostResponse[T]
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response.Result, nil
}
