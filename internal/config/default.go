package config

import (
	"time"

	"github.com/Idea-Thrive/backend/internal/http"
	"github.com/Idea-Thrive/backend/internal/jwt"
	"github.com/Idea-Thrive/backend/internal/logger"
	"github.com/Idea-Thrive/backend/internal/mysql"
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
		DB: mysql.Config{
			Host: "185.220.224.114",
			User: "root",
			Pass: "root",
			Port: "3306",
			Name: "app",
		},
	}
}
