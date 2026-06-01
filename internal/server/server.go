package server

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
	"context"

	"github.com/exceptnull/curl-imposter/internal/storage" // Update if your module name differs
)

// Server holds the mock data and manages the HTTP lifecycle
type Server struct {
	mu    sync.RWMutex
	mocks map[string]*storage.MockEntry
	port  int
	srv   *http.Server
}

// New creates and initializes a mock server, loading entries from disk
func New(port int, dataDir string) (*Server, error) {
	s := &Server{
		mocks: make(map[string]*storage.MockEntry),
		port:  port,
	}

	if err := s.loadMocks(dataDir); err != nil {
		return nil, err
	}

	return s, nil
}

// loadMocks reads stored entries and builds a fast lookup map: "METHOD /path" -> Mock
func (s *Server) loadMocks(dataDir string) error {
	entries, err := storage.Load(dataDir)
	if err != nil {
		return fmt.Errorf("failed to load mocks: %w", err)
	}

	for _, m := range entries {
		// Normalize URL to extract just the path for routing
		u, err := url.Parse(m.URL)
		if err != nil {
			fmt.Printf("⚠️ Skipping malformed URL: %s\n", m.URL)
			continue
		}
		
		// Map key: "GET /api/v1/users"
		key := strings.ToUpper(m.Method) + " " + u.Path
		s.mocks[key] = &m
	}

	return nil
}

// Start launches the mock HTTP server
func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleRequest)

	addr := fmt.Sprintf(":%d", s.port)
	s.srv = &http.Server{Addr: addr, Handler: mux}

	fmt.Printf("🚀 Mock server running on http://localhost:%d\n", s.port)
	s.printLoadedRoutes()
	
	return s.srv.ListenAndServe()
}

// handleRequest matches incoming requests to stored mocks and replays them
func (s *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Build lookup key from incoming request
	key := strings.ToUpper(r.Method) + " " + r.URL.Path
	mock, found := s.mocks[key]

	if !found {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf(
			`{"error":"Mock not found", "method":"%s", "path":"%s", "hint":"Run 'curl-imposter learn \"curl ...\"' to record it"}`,
			r.Method, r.URL.Path,
		)))
		return
	}

	// Simulate original API latency for realism
	if mock.LatencyMs > 0 {
		time.Sleep(time.Duration(mock.LatencyMs) * time.Millisecond)
	}

	// Write status code
	w.WriteHeader(mock.StatusCode)

	// Replay headers (skip auto-managed ones)
	for k, v := range mock.Headers {
		switch strings.ToLower(k) {
		case "content-length", "transfer-encoding", "connection":
			continue // Go handles these automatically
		default:
			w.Header().Set(k, v)
		}
	}

	// Write response body
	w.Write([]byte(mock.Body))
}

// printLoadedRoutes displays a nice startup summary
func (s *Server) printLoadedRoutes() {
	if len(s.mocks) == 0 {
		fmt.Println("📭 No mocks loaded yet. Use 'curl-imposter learn' to capture APIs.")
		return
	}

	fmt.Println("📊 Loaded endpoints:")
	for key := range s.mocks {
		fmt.Printf("   • %s\n", key)
	}
}

// Shutdown gracefully stops the HTTP server with a timeout
func (s *Server) Shutdown(ctx context.Context) error {
	if s.srv == nil {
		return nil
	}
	return s.srv.Shutdown(ctx)
}