package handler

import (
	"github.com/Idea-Thrive/backend/internal/model"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Category struct {
	Store  *store.Store
	Logger *zap.Logger
}

// Register function.
func (c Category) Register(group fiber.Router) {
	group.Post("/", c.Create)
}

func (c Category) Create(ctx *fiber.Ctx) error {
	category := new(model.Category)

	if err := ctx.BodyParser(category); err != nil {
		c.Logger.Error("failed to parse body", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.Store.CategoryCreate(*category); err != nil {
		c.Logger.Error("failed to create category", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusCreated)
}
