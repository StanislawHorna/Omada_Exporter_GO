package ClientInstrumentation

import (
	"regexp"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	label_clientType string = "clientType"
	label_method     string = "method"
	label_url        string = "url"
	label_code       string = "code"
)

var omadaHttpClientInstrumentationLabels = []string{label_clientType, label_method, label_url, label_code}

var (
	omada_http_client_requests_total = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "omada_http_client_requests_total",
			Help: "Total number of Omada HTTP client requests",
		},
		omadaHttpClientInstrumentationLabels,
	)
	omada_http_client_request_duration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "omada_http_client_request_duration_seconds",
			Help:    "Duration of HTTP client requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		omadaHttpClientInstrumentationLabels,
	)
)

// Precompiled regexes for performance
var (
	// UUIDs (with or without hyphens)
	uuidRegex = regexp.MustCompile(`(?i)[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}`)
	// MAC addresses with dashes or colons
	macRegex = regexp.MustCompile(`(?i)([0-9a-f]{2}[:-]){5}[0-9a-f]{2}`)
	// 24+ character hex strings
	hexRegex = regexp.MustCompile(`(?i)\b[0-9a-f]{24,}\b`)
	// Pure numeric IDs
	numericIDRegex = regexp.MustCompile(`^\d+$`)
)

func sanitizePath(path string) string {
	segments := strings.Split(path, "/")

	for i, segment := range segments {
		switch {
		case uuidRegex.MatchString(segment):
			segments[i] = ":uuid"
		case macRegex.MatchString(segment):
			segments[i] = ":mac"
		case hexRegex.MatchString(segment):
			segments[i] = ":hex"
		case numericIDRegex.MatchString(segment):
			segments[i] = ":id"
		}
	}

	return strings.Join(segments, "/")
}

func getHttpClientInstrumentationLabels(clientType OmadaClientType, method string, url string, code string) prometheus.Labels {
	return prometheus.Labels{
		label_clientType: string(clientType),
		label_method:     method,
		label_url:        sanitizePath(url),
		label_code:       code,
	}
}
