package api

import (
	"github.com/Idea-Thrive/backend/internal/api/user"
	"github.com/Idea-Thrive/backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// route function.
func route(app *fiber.App, handler *http.Handler) {
	userHandler := user.NewUserHandler(&model.UserModel{})

	app.Post("/create", userHandler.Create)
}
