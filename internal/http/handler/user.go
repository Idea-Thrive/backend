package handler

import (
	"github.com/Idea-Thrive/backend/internal/http/request"
	"github.com/Idea-Thrive/backend/internal/model"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// User struct.
type User struct {
	Store  *store.Store
	Logger *zap.Logger
}

// Register function.
func (usr User) Register(group fiber.Router) {
	group.Post("/users", usr.Create)
}

// Create function.
func (usr User) Create(ctx *fiber.Ctx) error {
	req := new(request.UserCreation)

	if err := ctx.BodyParser(req); err != nil {
		usr.Logger.Error("failed to parse request body", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{ //nolint:wrapcheck
			"message": err.Error(),
		})
	}

	user := model.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		PhotoURL:    req.PhotoURL,
		PersonnelID: req.PersonnelID,
		Gender:      req.Gender,
		Role:        req.Role,
	}

	if err := usr.Store.UserCreate(user); err != nil {
		usr.Logger.Error("failed to create user", zap.Error(err))

		return ctx.Status(fiber.StatusExpectationFailed).JSON(req) //nolint:wrapcheck
	}

	usr.Logger.Info("create user successfully")

	return ctx.Status(fiber.StatusOK).JSON(req) //nolint:wrapcheck
}
