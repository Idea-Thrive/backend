package http

import (
	"github.com/Idea-Thrive/backend/internal/http/handler"
	"github.com/Idea-Thrive/backend/internal/http/middleware"
	"github.com/Idea-Thrive/backend/internal/jwt"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// App struct.
type App struct {
	JWT    *jwt.JWT
	Store  *store.Store
	Logger *zap.Logger
}

// Register function.
func (a App) Register(app *fiber.App) {
	auth := middleware.Auth{
		JWT:    a.JWT,
		Logger: a.Logger,
		Store:  a.Store,
	}

	handler.Authentication{
		JWT:    a.JWT,
		Logger: a.Logger,
		Store:  a.Store,
	}.Register(app.Group("/auth"))

	handler.User{
		Store:  a.Store,
		Logger: a.Logger,
	}.Register(app.Group("/users", auth.Auth))

	handler.Idea{
		Store:  a.Store,
		Logger: a.Logger,
	}.Register(app.Group("/ideas", auth.Auth))

	handler.Company{
		Store:  a.Store,
		Logger: a.Logger,
	}.Register(app.Group("/companies", auth.Auth))

	handler.Category{
		Store:  a.Store,
		Logger: a.Logger,
	}.Register(app.Group("/categories", auth.Auth))

	handler.Criteria{
		Store:  a.Store,
		Logger: a.Logger,
	}.Register(app.Group("/criteria", auth.Auth))

	handler.Comment{
		Store:  a.Store,
		Logger: a.Logger,
	}.Register(app.Group("/comments", auth.Auth))
}
