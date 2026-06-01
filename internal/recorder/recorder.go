package recorder

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/exceptnull/curl-imposter/internal/parser" // Update if your module name differs
)

// RecordedResponse holds the complete API response captured from the real endpoint
type RecordedResponse struct {
	StatusCode int                 `json:"status_code"`
	Headers    map[string]string   `json:"headers"`
	Body       string              `json:"body"`
	LatencyMs  int64               `json:"latency_ms"`
}

// Record executes the parsed curl request against the real API and captures the response
func Record(req *parser.ParsedRequest) (*RecordedResponse, error) {
	// Prepare request body if present
	var bodyReader io.Reader
	if req.Body != "" {
		bodyReader = strings.NewReader(req.Body)
	}

	// Build the HTTP request
	httpReq, err := http.NewRequest(req.Method, req.URL, bodyReader)
	if err != nil {
		return nil, err
	}

	// Inject headers from the curl command
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	// Create client with a sensible timeout
	client := &http.Client{Timeout: 30 * time.Second}

	// Execute and measure latency
	start := time.Now()
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	latency := time.Since(start).Milliseconds()

	// Read the full response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Capture response headers (join multi-value headers with commas)
	capturedHeaders := make(map[string]string)
	for k, v := range resp.Header {
		if len(v) > 0 {
			capturedHeaders[k] = strings.Join(v, ", ")
		}
	}

	return &RecordedResponse{
		StatusCode: resp.StatusCode,
		Headers:    capturedHeaders,
		Body:       string(respBody),
		LatencyMs:  latency,
	}, nil
}