package ClientInstrumentation

import (
	"fmt"
	"net/http"
	"time"
)

type OmadaClientType string

const (
	OpenApiClientType OmadaClientType = "openApi"
	WebApiClientType  OmadaClientType = "webApi"
)

type InstrumentedRoundTripper struct {
	RoundTripper http.RoundTripper
	ClientType   OmadaClientType
}

func (irt *InstrumentedRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := irt.RoundTripper.RoundTrip(req)
	duration := time.Since(start).Seconds()

	code := "0"
	if resp != nil {
		code = fmt.Sprintf("%d", resp.StatusCode)
	}

	labels := getHttpClientInstrumentationLabels(
		irt.ClientType,
		req.Method,
		req.URL.Path,
		code,
	)

	omada_http_client_requests_total.With(labels).Inc()
	omada_http_client_request_duration.With(labels).Observe(duration)

	return resp, err
}
