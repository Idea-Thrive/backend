package cmd

import (
	"os"

	"github.com/Idea-Thrive/backend/internal/cmd/serve"
)

// Execute function.
func Execute() {
	//rootCmd := &cobra.Command{
	//	Use:   "backend",
	//	Short: "backend of idea-thrive service",
	//	Long:  `backend of idea-thrive service`,
	//}
	//
	//rootCmd.AddCommand(
	//	serve.Command(),
	//)

	err := serve.Command().Execute()
	if err != nil {
		os.Exit(1)
	}
}
