package cmd

import (
	"github.com/Idea-Thrive/backend/internal/cmd/serve"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "backend",
		Short: "backend of idea-thrive service",
		Long:  `backend of idea-thrive service`,
	}

	rootCmd.AddCommand(
		serve.Command(),
	)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
