package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "backend",
		Short: "backend of idea-thrive service",
		Long:  `backend of idea-thrive service`,
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
