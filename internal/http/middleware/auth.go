package middleware

import (
	"strings"

	"github.com/Idea-Thrive/backend/internal/jwt"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Auth struct.
type Auth struct {
	JWT    *jwt.JWT
	Store  *store.Store
	Logger *zap.Logger
}

// Auth function.
func (a Auth) Auth(ctx *fiber.Ctx) error {
	auth := ctx.Get("Authorization")

	if len(auth) == 0 {
		a.Logger.Error("no authorization header")

		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	token := strings.Split(auth, " ")[1]

	payload, err := a.JWT.Verify(token)
	if err != nil {
		a.Logger.Error("Failed to verify token", zap.Error(err))

		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Failed to verify token",
		})
	}

	if err := payload.Valid(); err != nil {
		a.Logger.Error("token is not valid", zap.Error(err))

		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Failed to verify token",
		})
	}

	ctx.Locals("username", payload.Email)

	a.Logger.Info("Successfully verified token", zap.String("username", payload.Email))

	return ctx.Next()
}
