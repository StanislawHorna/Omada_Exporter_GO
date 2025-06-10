package utils

import (
	"net/url"
	"strings"
)

func CreateURL(baseURL string, endpoint string, params map[string]string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	u.Path = u.Path + endpoint
	q := u.Query()
	if params != nil {
		for key, value := range params {
			q.Set(key, value)
		}
		u.RawQuery = q.Encode()
	}

	return u.String(), nil
}

func FillInEndpointPlaceholders(endpoint string, placeholders map[string]string) string {
	for key, value := range placeholders {
		placeholder := "{" + key + "}"
		endpoint = strings.Replace(endpoint, placeholder, value, -1)
	}
	return endpoint
}
