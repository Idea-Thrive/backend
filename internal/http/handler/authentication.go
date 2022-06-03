package handler

import (
	"net/http"

	"github.com/Idea-Thrive/backend/internal/http/request"
	"github.com/Idea-Thrive/backend/internal/jwt"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Authentication strcut.
type Authentication struct {
	JWT    jwt.JWT
	Store  *store.Store
	Logger *zap.Logger
}

// Register function.
func (a Authentication) Register(group fiber.Router) {
	group.Post("/login", a.login)
}

// login function.
func (a *Authentication) login(ctx *fiber.Ctx) error {
	a.Logger.Info("login request")

	var req request.Login
	if err := ctx.BodyParser(&req); err != nil {
		a.Logger.Error("failed to parse body", zap.Error(err))

		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{ //nolint:wrapcheck
			"error": "failed to parse body",
		})
	}

	isValid, err := a.Store.Login(req.Username, req.Password)
	if err != nil {
		a.Logger.Error("failed to login user", zap.Error(err))

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{ //nolint:wrapcheck
			"error": "failed to login user",
		})
	}

	if !isValid {
		a.Logger.Info("user is not valid", zap.String("username", req.Username))

		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{ //nolint:wrapcheck
			"error": "user is not valid",
		})
	}

	signedToken, expirationDate, err := a.JWT.Generate(req.Username)
	if err != nil {
		a.Logger.Error("failed to signed user signedToken", zap.Error(err))

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{ //nolint:wrapcheck
			"error": "failed to signed user signedToken",
		})
	}

	a.Logger.Info("user logged in successfully", zap.String("username", req.Username))

	return ctx.Status(http.StatusOK).JSON(fiber.Map{ //nolint:wrapcheck
		"token":      signedToken,
		"expiration": expirationDate,
	})
}
