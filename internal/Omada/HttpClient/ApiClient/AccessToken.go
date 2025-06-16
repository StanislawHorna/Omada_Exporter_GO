package ApiClient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	utils "omada_exporter_go/internal/Omada/HttpClient/Utils"
)

const PATH_REQUEST_ACCESS_TOKEN = "/openapi/authorize/token"

type OpenApiTokenPayload struct {
	OmadaID      string `json:"omadacId"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type OpenApiAccessToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
	ExpiresIn    int    `json:"expiresIn"`
}

func (OpenApiAccessToken) Path(map[string]string) string {
	return PATH_REQUEST_ACCESS_TOKEN
}

func (OpenApiAccessToken) Payload(data map[string]any) (any, error) {
	var payload OpenApiTokenPayload
	utils.MapToStruct(data, &payload)
	return payload, nil
}

type AccessToken struct {
	response       *OpenApiAccessToken
	clientID       string
	clientSecret   string
	omadaID        string
	BaseURL        string
	httpClient     *http.Client
	expirationDate int64
}

func NewAccessToken(baseURL string, payload OpenApiTokenPayload) (*AccessToken, error) {
	if payload.ClientID == "" || payload.ClientSecret == "" || payload.OmadaID == "" {
		return nil, fmt.Errorf("missing required fields in OpenApiRequestToken: ClientID, ClientSecret, or OmadaID")
	}
	var a AccessToken
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	a.httpClient = &http.Client{Transport: customTransport}
	a.clientID = payload.ClientID
	a.clientSecret = payload.ClientSecret
	a.omadaID = payload.OmadaID
	a.BaseURL = baseURL

	if err := a.requestAccessToken(payload); err != nil {
		return nil, fmt.Errorf("failed to request access token: %w", err)
	}

	return &a, nil
}

func (a *AccessToken) requestAccessToken(payload OpenApiTokenPayload) error {
	url, err := utils.CreateURL(
		a.BaseURL,
		PATH_REQUEST_ACCESS_TOKEN,
		map[string]string{
			"grant_type": "client_credentials",
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create URL: %w", err)
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	response, err := a.httpClient.Post(url, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		return fmt.Errorf("failed to request access token: %w", err)
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get access token, status code: %d", response.StatusCode)
	}

	var omadaResult Response[OpenApiAccessToken]
	if err := json.NewDecoder(response.Body).Decode(&omadaResult); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	a.response = &omadaResult.Result
	a.expirationDate = time.Now().Unix() + int64(a.response.ExpiresIn)

	return nil
}

func (a *AccessToken) GetAccessToken() (string, error) {
	if a.response == nil {
		return "", fmt.Errorf("access token response is nil, please request a token first")
	}

	// Check if the token is about to expire in the next 5 minutes (300 seconds)
	if time.Now().Unix() >= (a.expirationDate - 300) {
		if err := a.requestAccessToken(OpenApiTokenPayload{
			OmadaID:      a.omadaID,
			ClientID:     a.clientID,
			ClientSecret: a.clientSecret,
		}); err != nil {
			return "", fmt.Errorf("failed to refresh access token: %w", err)
		}
	}

	return a.response.AccessToken, nil
}
