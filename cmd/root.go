package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// DataDir is the global directory where mock data is stored
	DataDir string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "curl-imposter",
	Short: "Mock any API instantly from a single curl command.",
	Long: `curl-imposter is a zero-config CLI tool that learns from your curl commands 
and spins up a local mock server to replay the exact responses. 

Perfect for offline development, testing, and sharing API behavior.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Determine default data directory (~/.curl-imposter)
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Warning: Could not determine home directory, using current directory.")
		home = "."
	}
	defaultDataDir := fmt.Sprintf("%s/.curl-imposter", home)

	// Add a persistent flag for the data directory (available to all subcommands)
	rootCmd.PersistentFlags().StringVar(&DataDir, "data-dir", defaultDataDir, "Directory to store mock data")
}