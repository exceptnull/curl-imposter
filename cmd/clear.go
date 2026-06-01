package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/exceptnull/curl-imposter/internal/storage" // Update if your module name differs
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Remove all stored mock data",
	Long: `Deletes the mocks database file from the data directory.
This resets the mock server to an empty state.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if file exists first to give a friendly message
		mocksFile := filepath.Join(DataDir, "mocks.json")
		if _, err := os.Stat(mocksFile); os.IsNotExist(err) {
			fmt.Println("📭 No mock data found. Already clean!")
			return nil
		}

		if err := storage.Clear(DataDir); err != nil {
			return fmt.Errorf("❌ failed to clear mocks: %w", err)
		}

		fmt.Println("🧹 Mock database cleared successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}