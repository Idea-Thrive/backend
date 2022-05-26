package handler

import (
	"github.com/Idea-Thrive/backend/internal/http/request"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const TokeExpiration = 60 * 24 * time.Minute

type Authentication struct {
	Secret string

	Store  *store.Store
	Logger *zap.Logger
}

func (a Authentication) Register(group fiber.Router) {
	group.Post("/login", a.login)
}

func (a *Authentication) login(ctx *fiber.Ctx) error {

	var req request.Login
	if err := ctx.BodyParser(&req); err != nil {
		a.Logger.Error("failed to parse body", zap.Error(err))

		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse body",
		})
	}

	isValid, err := a.Store.Login(req.Username, req.Password)
	if err != nil {
		a.Logger.Error("failed to login user", zap.Error(err))

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to login user",
		})
	}

	if !isValid {
		a.Logger.Info("user is not valid", zap.String("username", req.Username))

		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"error": "user is not valid",
		})
	}

	expirationDate := time.Now().Add(TokeExpiration).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": req.Username,
		"exp":     expirationDate,
	})

	signedToken, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		a.Logger.Error("failed to signed user token", zap.Error(err))

		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to signed user token",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"token":      signedToken,
		"expiration": expirationDate,
	})
}
