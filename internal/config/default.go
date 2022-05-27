package config

import (
	"github.com/Idea-Thrive/backend/internal/http"
	"github.com/Idea-Thrive/backend/internal/logger"
	"github.com/Idea-Thrive/backend/internal/mysql"
)

// Default function.
func Default() Config {
	return Config{
		HTTP: http.Config{
			Port:   8080,
			Secret: "jafar",
		},
		Log: logger.Config{Level: "debug"},
		DB: mysql.Config{
			Host: "localhost",
			User: "root",
			Pass: "amir_hallaji",
			Port: "3306",
			Name: "app",
		},
	}
}
