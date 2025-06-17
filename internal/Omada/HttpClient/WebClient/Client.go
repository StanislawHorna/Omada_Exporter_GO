package WebClient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"sync"

	"omada_exporter_go/internal"
	"omada_exporter_go/internal/Omada/HttpClient/ApiClient"
	"omada_exporter_go/internal/Omada/HttpClient/Utils"
)

const (
	path_login        = "/{omadaID}/api/v2/login"
	path_logout       = "/{omadaID}/api/v2/logout"
	path_login_status = "/{omadaID}/api/v2/loginStatus"
)

type WebClient struct {
	BaseURL  string
	OmadaID  string
	username string
	password string
	SiteID   string
	SiteName string
	Client   *http.Client
	Token    string
}

func (w *WebClient) fillInOmadaIDs(placeholders map[string]string) map[string]string {
	if placeholders == nil {
		placeholders = make(map[string]string)
	}
	placeholders["omadaID"] = w.OmadaID
	placeholders["siteID"] = w.SiteID
	return placeholders
}

var (
	instance *WebClient
	once     sync.Once
)

func newClient(baseURL string, username string, password string, siteName string) *WebClient {
	jar, _ := cookiejar.New(nil)
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	openApiClient := ApiClient.GetInstance()

	clientObject := &WebClient{
		BaseURL:  baseURL,
		OmadaID:  openApiClient.OmadaID,
		username: username,
		password: password,
		SiteName: siteName,
		SiteID:   openApiClient.SiteID,

		Client: &http.Client{
			Jar:       jar,
			Transport: customTransport,
		},
	}
	clientObject.Login()

	if !clientObject.isLoggedIn() {
		fmt.Println("Failed to log in to Omada controller")
		return nil
	}

	return clientObject
}

func GetInstance() *WebClient {
	once.Do(func() {
		conf := internal.GetConfig().Omada
		instance = newClient(conf.OmadaURL, conf.Username, conf.Password, conf.SiteName)
	})
	return instance
}

func (c *WebClient) Login() error {
	endpoint := Utils.FillInEndpointPlaceholders(path_login, c.fillInOmadaIDs(nil))
	if endpoint == "" {
		fmt.Println("Endpoint cannot be empty")
		return fmt.Errorf("endpoint cannot be empty")
	}

	url, err := Utils.CreateURL(c.BaseURL, endpoint, nil)
	if err != nil {
		return fmt.Errorf("error creating URL: %w", err)
	}
	bodyBytes, err := json.Marshal(map[string]string{"username": c.username, "password": c.password})
	if err != nil {
		return fmt.Errorf("error marshalling request body: %w", err)
	}

	response, err := c.Client.Post(url, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("error making POST request: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("non-ok status code: %d", response.StatusCode)
	}

	defer response.Body.Close()
	var loginResponse Response[Login]
	if err := json.NewDecoder(response.Body).Decode(&loginResponse); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}
	if loginResponse.ErrorCode != 0 {
		return fmt.Errorf("API error: %s (code %d)", loginResponse.Message, loginResponse.ErrorCode)
	}
	c.Token = loginResponse.Result.Token
	return nil
}

func (c *WebClient) isLoggedIn() bool {
	endpoint := Utils.FillInEndpointPlaceholders(path_login_status, c.fillInOmadaIDs(nil))
	if endpoint == "" {
		fmt.Println("Endpoint cannot be empty")
		return false
	}

	url, err := Utils.CreateURL(c.BaseURL, endpoint, Utils.AddTimestampParam(nil))
	if err != nil {
		fmt.Println("Error creating URL:", err)
		return false
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return false
	}

	if err := c.setAuthorizationHeader(req); err != nil {
		fmt.Println("Error setting authorization header:", err)
		return false
	}

	response, err := c.Client.Do(req)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return false
	}
	defer response.Body.Close()
	var result Response[IsLoggedIn]
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return false
	}
	if result.ErrorCode != 0 {
		return false
	}

	return result.Result.Login
}

func (c *WebClient) setAuthorizationHeader(req *http.Request) error {
	req.Header.Set("Csrf-Token", c.Token)
	return nil
}
