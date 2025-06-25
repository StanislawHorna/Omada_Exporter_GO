package Log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type LokiLabels struct {
	env        string
	appName    string
	appVersion string
	goVersion  string
}

func (ll LokiLabels) toMap() map[string]string {
	return map[string]string{
		"env":             ll.env,
		"service_name":    ll.appName,
		"service_version": ll.appVersion,
		"go_version":      ll.goVersion,
	}
}

type LokiEntry struct {
	Stream map[string]string `json:"stream"`
	Values [][2]string       `json:"values"`
}

type LokiPayload struct {
	Streams []LokiEntry `json:"streams"`
}

type LokiWriter struct {
	endpoint string
	labels   map[string]string
	client   *http.Client
	mu       sync.Mutex
}

func newLokiWriter(endpoint string, labels LokiLabels) *LokiWriter {
	return &LokiWriter{
		endpoint: endpoint,
		labels:   labels.toMap(),
		client:   &http.Client{Timeout: 5 * time.Second},
	}
}

func (lw *LokiWriter) Write(p []byte) (n int, err error) {
	lw.mu.Lock()
	defer lw.mu.Unlock()

	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)

	payload := LokiPayload{
		Streams: []LokiEntry{
			{
				Stream: lw.labels,
				Values: [][2]string{
					{timestamp, string(bytes.TrimSpace(p))},
				},
			},
		},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", lw.endpoint+"/loki/api/v1/push", bytes.NewBuffer(data))
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := lw.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return 0, fmt.Errorf("unexpected status code from Loki: %d", resp.StatusCode)
	}

	return len(p), nil
}
