package handler

import (
	"github.com/Idea-Thrive/backend/internal/model"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Criteria struct {
	Store  *store.Store
	Logger *zap.Logger
}

// Register function.
func (c Criteria) Register(group fiber.Router) {
	group.Post("/", c.Create)
	group.Get("/", c.GetAll)
	group.Delete("/:id", c.Delete)
}

func (c Criteria) Create(ctx *fiber.Ctx) error {
	criteria := new(model.Criteria)

	if err := ctx.BodyParser(&criteria); err != nil {
		c.Logger.Error("failed to parse body", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.Store.CriteriaCreate(*criteria); err != nil {
		c.Logger.Error("failed to create criteria", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (c Criteria) GetAll(ctx *fiber.Ctx) error {
	categoryID := ctx.Query("category_id")

	criteria, err := c.Store.CriteriaGetAll(categoryID)
	if err != nil {
		c.Logger.Error("failed to get criteria", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(criteria)
}

func (c Criteria) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.Store.CriteriaDelete(id); err != nil {
		c.Logger.Error("failed to delete criteria", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
