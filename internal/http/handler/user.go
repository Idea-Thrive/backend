package handler

import (
	"strconv"

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
func (u User) Register(group fiber.Router) {
	group.Post("/", u.Create)
	group.Get("/", u.GetAll)
	group.Get("/info", u.GetByUsername)
	group.Get("/:id", u.Get)
	group.Put("/:id", u.Update)
	group.Delete("/:id", u.Delete)
}

// Create function.
func (u User) Create(ctx *fiber.Ctx) error {
	req := new(request.UserCreation)

	if err := ctx.BodyParser(req); err != nil {
		u.Logger.Error("failed to parse request body", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := model.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		PhotoURL:    req.PhotoURL,
		CompanyID:   req.CompanyID,
		PersonnelID: req.PersonnelID,
		Gender:      req.Gender,
		Role:        req.Role,
	}

	if err := u.Store.UserCreate(user); err != nil {
		u.Logger.Error("failed to create user", zap.Error(err))

		return ctx.Status(fiber.StatusExpectationFailed).JSON(req)
	}

	u.Logger.Info("user created successfully")

	return ctx.Status(fiber.StatusOK).JSON(req)
}

// Get function.
func (u User) Get(ctx *fiber.Ctx) error {
	userID := ctx.AllParams()["id"]

	user, err := u.Store.UserGet(userID)
	if err != nil {
		u.Logger.Error("failed to get user from store", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

// GetByUsername function.
func (u User) GetByUsername(ctx *fiber.Ctx) error {
	username := ctx.Locals("username").(string)

	user, err := u.Store.UserGetByUsername(username)
	if err != nil {
		u.Logger.Error("failed to get user from store", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

// Update function.
func (u User) Update(ctx *fiber.Ctx) error {
	userID := ctx.AllParams()["id"]

	req := new(request.UserCreation)

	if err := ctx.BodyParser(req); err != nil {
		u.Logger.Error("failed to parse request body", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user := model.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		PhotoURL:    req.PhotoURL,
		CompanyID:   req.CompanyID,
		PersonnelID: req.PersonnelID,
		Gender:      req.Gender,
		Role:        req.Role,
	}

	if err := u.Store.UserUpdate(userID, user); err != nil {
		u.Logger.Error("failed to update user", zap.Error(err))

		return ctx.Status(fiber.StatusExpectationFailed).JSON(req)
	}

	u.Logger.Info("user updated successfully")

	return ctx.Status(fiber.StatusOK).JSON(req)
}

// Delete function.
func (u User) Delete(ctx *fiber.Ctx) error {
	userID := ctx.AllParams()["id"]

	if err := u.Store.UserDelete(userID); err != nil {
		u.Logger.Error("failed to delete user")

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// GetAll function.
func (u User) GetAll(ctx *fiber.Ctx) error {
	size, _ := strconv.Atoi(ctx.Query("size", "100"))   // optional
	offset, _ := strconv.Atoi(ctx.Query("offset", "0")) // optional

	users, err := u.Store.UserGetAll(size, offset)
	if err != nil {
		u.Logger.Error("failed to get users from store", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	u.Logger.Info("users retrieved successfully")

	return ctx.Status(fiber.StatusOK).JSON(users)
}
