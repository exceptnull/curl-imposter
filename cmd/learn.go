package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/exceptnull/curl-imposter/internal/parser"
	"github.com/exceptnull/curl-imposter/internal/recorder"
	"github.com/exceptnull/curl-imposter/internal/storage"
)

var learnCmd = &cobra.Command{
	Use:   "learn [curl-command]",
	Short: "Capture a real API response to use as a mock",
	Long: `Executes the provided curl command, captures the full response 
(status, headers, body, latency), and saves it locally for mocking.

Example:
  curl-imposter learn "curl -X POST https://api.example.com/users -H 'Content-Type: application/json' -d '{"name":"test"}'"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		curlString := args[0]

		// 1️⃣ Parse the curl command
		parsed, err := parser.Parse(curlString)
		if err != nil {
			return fmt.Errorf("❌ failed to parse curl command: %w", err)
		}
		fmt.Printf("🔍 Parsed: %s %s\n", parsed.Method, parsed.URL)

		// 2️⃣ Hit the real API and record the response
		fmt.Println("📡 Hitting real API...")
		response, err := recorder.Record(parsed)
		if err != nil {
			return fmt.Errorf("❌ failed to record response: %w", err)
		}
		fmt.Printf("✅ Captured: %d (%dms)\n", response.StatusCode, response.LatencyMs)

		// 3️⃣ Save to disk as a mock entry
		mockEntry := storage.MockEntry{
			Method:         parsed.Method,
			URL:            parsed.URL,
			RequestHeaders: parsed.Headers,
			StatusCode:     response.StatusCode,
			Headers:        response.Headers,
			Body:           response.Body,
			LatencyMs:      response.LatencyMs,
		}

		if err := storage.AddOrUpdate(DataDir, mockEntry); err != nil {
			return fmt.Errorf("❌ failed to save mock: %w", err)
		}

		fmt.Printf("💾 Mock saved to %s/mocks.json\n", DataDir)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(learnCmd)
}