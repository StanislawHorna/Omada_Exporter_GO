package ApiClient

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"omada_exporter_go/internal/Omada/HttpClient/Requests"
	Utils "omada_exporter_go/internal/Omada/HttpClient/Utils"
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

func (c *ApiClient) getApiInfo() (*Requests.ApiInfoResponse, error) {
	if c.Http == nil {
		return nil, fmt.Errorf("HTTP client is not initialized")
	}
	url, err := Utils.CreateURL(c.BaseURL, API_INFO_PATH, nil)
	if err != nil {
		fmt.Println("Error creating URL:", err)
		return nil, err
	}

	res, err := c.Http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, err
	}

	defer res.Body.Close()
	var apiInfoResponse Response[Requests.ApiInfoResponse]
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d from API\n", res.StatusCode)
		return nil, err
	}
	if err := json.NewDecoder(res.Body).Decode(&apiInfoResponse); err != nil {
		fmt.Println("Error decoding response:", err)
		return nil, err
	}

	c.OmadaID = apiInfoResponse.Result.OmadaID
	return &apiInfoResponse.Result, nil
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

	var err error

	_, err = apiClientObject.getApiInfo()
	if err != nil {
		fmt.Println("Error fetching API info:", err)
		return nil
	}

	apiClientObject.auth, err = NewAccessToken(
		apiClientObject.BaseURL,
		OpenApiTokenPayload{
			OmadaID:      apiClientObject.OmadaID,
			ClientID:     ClientID,
			ClientSecret: ClientSecret,
		},
	)
	if err != nil {
		fmt.Println("Error creating access token:", err)
		return nil
	}

	endpoint := Utils.FillInEndpointPlaceholders(Requests.PATH_SITES, map[string]string{"omadaID": apiClientObject.OmadaID})

	res, err := Get[Requests.Sites](*apiClientObject, endpoint, map[string]string{"omadaID": apiClientObject.OmadaID}, nil)

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
