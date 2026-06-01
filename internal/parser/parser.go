package parser

import (
	"errors"
	"strings"

	"github.com/kballard/go-shellquote"
)

// ParsedRequest holds the structured data extracted from a curl command
type ParsedRequest struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    string
}

// Parse takes a raw curl command string and converts it into a ParsedRequest
func Parse(curlCmd string) (*ParsedRequest, error) {
	// shellquote.Split safely handles nested quotes like: -d '{"key": "value"}'
	args, err := shellquote.Split(curlCmd)
	if err != nil {
		return nil, err
	}

	req := &ParsedRequest{
		Method:  "GET", // Default method
		Headers: make(map[string]string),
	}

	expecting := "" // Tracks what the next argument should be (e.g., "header", "body")

	for _, arg := range args {
		// If we are expecting a value for a flag, process it
		if expecting != "" {
			switch expecting {
			case "method":
				req.Method = strings.ToUpper(arg)
			case "header":
				parseHeader(req, arg)
			case "body":
				req.Body = arg
				// If a body is provided and method is still GET, curl defaults to POST
				if req.Method == "GET" {
					req.Method = "POST"
				}
			}
			expecting = ""
			continue
		}

		// Process flags
		switch arg {
		case "curl":
			continue // Skip the word "curl"
		case "-X", "--request":
			expecting = "method"
		case "-H", "--header":
			expecting = "header"
		case "-d", "--data", "--data-raw", "--data-binary":
			expecting = "body"
		case "-I", "--head":
			req.Method = "HEAD"
		default:
			// If it doesn't start with '-', and we aren't expecting a value, it must be the URL
			if !strings.HasPrefix(arg, "-") {
				req.URL = arg
			}
			// Ignore other unknown flags to prevent crashing on complex copy-pastes
		}
	}

	if req.URL == "" {
		return nil, errors.New("no URL found in curl command")
	}

	return req, nil
}

// parseHeader splits "Content-Type: application/json" into key/value and adds to map
func parseHeader(req *ParsedRequest, headerStr string) {
	parts := strings.SplitN(headerStr, ":", 2)
	if len(parts) == 2 {
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		req.Headers[key] = value
	}
}