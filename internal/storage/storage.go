package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// MockEntry represents a single recorded endpoint ready to be served as a mock
type MockEntry struct {
	Method         string            `json:"method"`
	URL            string            `json:"url"`               // Full original URL
	RequestHeaders map[string]string `json:"request_headers,omitempty"`
	StatusCode     int               `json:"status_code"`
	Headers        map[string]string `json:"response_headers"`
	Body           string            `json:"body"`
	LatencyMs      int64             `json:"latency_ms"`
}

const mocksFilename = "mocks.json"

// EnsureDir creates the data directory if it doesn't exist
func EnsureDir(dataDir string) error {
	return os.MkdirAll(dataDir, 0755)
}

func getFilePath(dataDir string) string {
	return filepath.Join(dataDir, mocksFilename)
}

// Load reads all mock entries from disk. Returns an empty slice if none exist.
func Load(dataDir string) ([]MockEntry, error) {
	path := getFilePath(dataDir)
	
	// If file doesn't exist yet, return empty slice (not an error)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []MockEntry{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Handle empty file gracefully
	if len(data) == 0 {
		return []MockEntry{}, nil
	}

	var mocks []MockEntry
	if err := json.Unmarshal(data, &mocks); err != nil {
		return nil, err
	}

	return mocks, nil
}

// Save writes the full mock list to disk with pretty formatting
func Save(dataDir string, mocks []MockEntry) error {
	if err := EnsureDir(dataDir); err != nil {
		return err
	}

	data, err := json.MarshalIndent(mocks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(getFilePath(dataDir), data, 0644)
}

// AddOrUpdate appends a new mock entry, or updates it if Method+URL already exists
func AddOrUpdate(dataDir string, newMock MockEntry) error {
	mocks, err := Load(dataDir)
	if err != nil {
		return err
	}

	// Check for existing entry with same method & URL
	for i, m := range mocks {
		if m.URL == newMock.URL && m.Method == newMock.Method {
			mocks[i] = newMock
			return Save(dataDir, mocks)
		}
	}

	// Append as new entry
	mocks = append(mocks, newMock)
	return Save(dataDir, mocks)
}

// Clear removes the mock data file if it exists
func Clear(dataDir string) error {
	path := getFilePath(dataDir)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}
	return os.Remove(path)
}