package ApiClient

import (
	"crypto/tls"
	"fmt"
	"net/http"
	OmadaRequests "omada_exporter_go/internal/Omada/HttpClient/Requests"
	response "omada_exporter_go/internal/Omada/HttpClient/Requests/GenericResponse"
	"sync"
)

const API_INFO_PATH = "/api/info"

type ApiClient struct {
	BaseURL  string
	OmadaID  string
	SiteID   string
	SiteName string
	Http     *http.Client
	auth     *AccessToken
}

var (
	instance *ApiClient
	once     sync.Once
)

func newClient(BaseURL string, ClientID string, ClientSecret string, SiteName string) *ApiClient {
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	apiClientObject := &ApiClient{
		BaseURL:  BaseURL,
		SiteName: SiteName,
		Http:     &http.Client{Transport: customTransport},
	}

	apiInfoResponse, err := response.Get[OmadaRequests.ApiInfoResponse](apiClientObject.Http, apiClientObject.BaseURL, nil, nil)

	if err != nil {
		fmt.Println("Error fetching API info:", err)
		return nil
	}

	apiClientObject.OmadaID = apiInfoResponse.OmadaID

	apiClientObject.auth, err = NewAccessToken(
		apiClientObject.BaseURL,
		OpenApiRequestToken{
			OmadaID:      apiClientObject.OmadaID,
			ClientID:     ClientID,
			ClientSecret: ClientSecret,
		},
	)
	if err != nil {
		fmt.Println("Error creating access token:", err)
		return nil
	}

	return apiClientObject
}

func GetInstance(BaseURL string, ClientID string, ClientSecret string, SiteID string) *ApiClient {
	once.Do(func() {
		instance = newClient(BaseURL, ClientID, ClientSecret, SiteID)
	})
	return instance
}
