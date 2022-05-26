package handler

import (
	"github.com/Idea-Thrive/backend/internal/http/request"
	"github.com/Idea-Thrive/backend/internal/model"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type User struct {
	Store  *store.Store
	Logger *zap.Logger
}

func (u User) Register(group fiber.Router) {
	group.Post("/users", u.create)
}

func (u User) create(c *fiber.Ctx) error {
	req := new(request.UserCreation)

	if err := c.BodyParser(req); err != nil {
		u.Logger.Error("failed to parse request body", zap.Error(err))

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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

	if err := u.Store.UserCreate(user); err != nil {
		u.Logger.Error("failed to create user", zap.Error(err))

		return c.Status(fiber.StatusExpectationFailed).JSON(req)
	}

	u.Logger.Info("create user successfully")

	return c.Status(fiber.StatusOK).JSON(req)
}
