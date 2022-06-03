package middleware

import (
	"strings"

	"github.com/Idea-Thrive/backend/internal/jwt"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Auth struct {
	JWT    *jwt.JWT
	Store  *store.Store
	Logger *zap.Logger
}

func (a Auth) Auth(ctx *fiber.Ctx) error {
	token := strings.Split(ctx.Get("Authorization"), " ")[1]

	payload, err := a.JWT.Verify(token)
	if err != nil {
		a.Logger.Error("Failed to verify token", zap.Error(err))

		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{ //nolint:wrapcheck
			"message": "Failed to verify token",
		})
	}

	ctx.Locals("username", payload.Username)

	a.Logger.Info("Successfully verified token", zap.String("username", payload.Username))

	return ctx.Next() //nolint:wrapcheck
}
