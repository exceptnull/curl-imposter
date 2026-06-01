package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/exceptnull/curl-imposter/internal/server" // Update if your module name differs
)

var port int

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the local mock server",
	Long: `Starts an HTTP server on the specified port that replays captured API responses.
Press Ctrl+C to stop the server gracefully.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Initialize the mock server
		mockSrv, err := server.New(port, DataDir)
		if err != nil {
			return fmt.Errorf("❌ failed to initialize mock server: %w", err)
		}

		// Start server in a background goroutine
		serverErr := make(chan error, 1)
		go func() {
			serverErr <- mockSrv.Start()
		}()

		// Wait for interrupt signal (Ctrl+C / SIGTERM)
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("\n👋 Shutting down mock server gracefully...")

		// Graceful shutdown with a 5-second timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := mockSrv.Shutdown(ctx); err != nil {
			return fmt.Errorf("❌ server forced to shutdown: %w", err)
		}

		// Check if the server exited with an unexpected error
		select {
		case err := <-serverErr:
			if err != nil && err != http.ErrServerClosed {
				return err
			}
		default:
		}

		fmt.Println("✅ Server stopped successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the mock server on")
}