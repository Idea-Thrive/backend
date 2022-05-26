package config

import (
	"github.com/Idea-Thrive/backend/internal/http"
	"github.com/Idea-Thrive/backend/internal/logger"
)

// Default function.
func Default() Config {
	return Config{
		HTTP: http.Config{
			Port:   8080,
			Secret: "jafar",
		},
		Log: logger.Config{Level: "debug"},
	}
}
