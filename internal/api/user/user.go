package user

import (
	"github.com/Idea-Thrive/backend/internal/http/request"
	"github.com/Idea-Thrive/backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserHandler model.UserImpl
}

func NewUserHandler(handler model.UserImpl) *UserHandler {
	return &UserHandler{
		UserHandler: handler,
	}
}

func (u *UserHandler) Create(c *fiber.Ctx) error {
	req := new(request.UserCreation)

	errParse := c.BodyParser(req)

	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errParse.Error(),
		})
	}

	user := model.UserModel{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		PhotoURL:    req.PhotoURL,
		PersonnelID: req.PersonnelID,
		Gender:      req.Gender,
		Role:        req.Role,
	}

	if err := u.UserHandler.Create(user); err != nil {
		return c.Status(fiber.StatusExpectationFailed).JSON(req)
	}

	return c.Status(fiber.StatusOK).JSON(req)
}
