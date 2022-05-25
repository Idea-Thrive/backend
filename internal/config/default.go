package config

import (
	"github.com/Idea-Thrive/backend/internal/http"
	"github.com/Idea-Thrive/backend/internal/logger"
)

func Default() Config {
	return Config{
		Http: http.Config{
			Port:   8080,
			Secret: "jafar"},
		Log: logger.Config{Level: "debug"},
	}
}
