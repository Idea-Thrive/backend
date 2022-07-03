package cmd

import (
	"os"

	"github.com/Idea-Thrive/backend/internal/cmd/serve"
)

// Execute function.
func Execute() {
	err := serve.Command().Execute()
	if err != nil {
		os.Exit(1)
	}
}
