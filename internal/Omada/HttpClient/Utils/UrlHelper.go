package Utils

import (
	"fmt"
	"net/url"
	"strings"
	"time"
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

func AddTimestampParam(params map[string]string) map[string]string {
	if params == nil {
		params = make(map[string]string)
	}
	params["_t"] = fmt.Sprintf("%d", time.Now().UnixMilli())
	return params
}

func FillInEndpointPlaceholders(endpoint string, placeholders map[string]string) string {
	for key, value := range placeholders {
		placeholder := "{" + key + "}"
		endpoint = strings.ReplaceAll(endpoint, placeholder, value)
	}
	return endpoint
}
