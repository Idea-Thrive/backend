package cmd

import (
	"os"

	"github.com/Idea-Thrive/backend/internal/cmd/serve"

	"github.com/spf13/cobra"
)

// Execute function.
func Execute() {
	rootCmd := &cobra.Command{
		Use:   "backend",
		Short: "backend of idea-thrive service",
		Long:  `backend of idea-thrive service`,
	}

	rootCmd.AddCommand(
		serve.Command(),
	)

	err := rootCmd.Execute() //nolint:ifshort
	if err != nil {
		os.Exit(1)
	}
}
