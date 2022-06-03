package config

import (
	"time"

	"github.com/Idea-Thrive/backend/internal/http"
	"github.com/Idea-Thrive/backend/internal/jwt"
	"github.com/Idea-Thrive/backend/internal/logger"
)

// Default function.
func Default() Config {
	return Config{
		HTTP: http.Config{
			Port: 8080,
		},
		JWT: jwt.Config{
			Expiration: 60 * 24 * time.Minute,
			Secret:     "jafar",
		},
		Log: logger.Config{Level: "debug"},
	}
}
