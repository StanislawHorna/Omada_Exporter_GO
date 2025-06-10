package ApiClient

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	OmadaRequests "omada_exporter_go/internal/Omada/HttpClient/Requests"
	response "omada_exporter_go/internal/Omada/HttpClient/Requests/GenericResponse"
	utils "omada_exporter_go/internal/Omada/HttpClient/Utils"
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

func (c *ApiClient) setAuthorizationHeader(req *http.Request) error {
	token, err := c.auth.GetAccessToken()
	if err != nil {
		return fmt.Errorf("error getting access token: %s", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("AccessToken=%s", token))
	return nil
}
func (c *ApiClient) fillInOmadaIDs(placeholders map[string]string) map[string]string {
	if placeholders == nil {
		placeholders = make(map[string]string)
	}
	placeholders["omadaID"] = c.OmadaID
	placeholders["siteID"] = c.SiteID
	return placeholders
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

	endpoint := OmadaRequests.SitesResponse{}.Path(map[string]string{"omadaID": apiClientObject.OmadaID})

	res, err := Get[OmadaRequests.SitesResponse](*apiClientObject, endpoint, map[string]string{"omadaID": apiClientObject.OmadaID}, nil)

	if err != nil {
		fmt.Println("Error fetching sites:", err)
		return nil
	}

	for _, site := range *res {
		if site.Name == apiClientObject.SiteName {
			apiClientObject.SiteID = site.SiteID
			break
		}
	}

	return apiClientObject
}

func GetInstance(BaseURL string, ClientID string, ClientSecret string, SiteID string) *ApiClient {
	once.Do(func() {
		instance = newClient(BaseURL, ClientID, ClientSecret, SiteID)
	})
	return instance
}

func Get[T any](client ApiClient, endpoint string, endpointPlaceholders map[string]string, queryParams map[string]string) (*[]T, error) {
	endpointPlaceholders = client.fillInOmadaIDs(endpointPlaceholders)
	endpoint = utils.FillInEndpointPlaceholders(endpoint, endpointPlaceholders)
	if endpoint == "" {
		fmt.Println("Endpoint cannot be empty")
		return nil, fmt.Errorf("endpoint cannot be empty")
	}

	var allData []T
	currentPage := 1

	for {
		queryParamsWithPage := AddPaginationParams(queryParams, currentPage)
		url, err := utils.CreateURL(client.BaseURL, endpoint, queryParamsWithPage)
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
