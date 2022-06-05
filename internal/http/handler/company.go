package handler

import (
	"time"

	"github.com/Idea-Thrive/backend/internal/model"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Company struct {
	Store  *store.Store
	Logger *zap.Logger
}

// Register function.
func (c Company) Register(group fiber.Router) {
	group.Post("/", c.Create)
	group.Get("/:id", c.Get)
	group.Delete("/:id", c.Delete)
}

func (c Company) Create(ctx *fiber.Ctx) error {
	company := new(model.Company)

	if err := ctx.BodyParser(company); err != nil {
		c.Logger.Error("failed to parse company", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{ //nolint:wrapcheck
			"error": err.Error(),
		})
	}

	company.CreatedAt = time.Now()
	company.UpdatedAt = time.Now()

	if err := c.Store.CompanyCreate(*company); err != nil {
		c.Logger.Error("failed to create company", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ //nolint:wrapcheck
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusOK) //nolint:wrapcheck
}

func (c Company) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	company, err := c.Store.CompanyGet(id)
	if err != nil {
		c.Logger.Error("failed to get company", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ //nolint:wrapcheck
			"error": err.Error(),
		})
	}

	return ctx.JSON(company) //nolint:wrapcheck
}

func (c Company) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.Store.CompanyDelete(id); err != nil {
		c.Logger.Error("failed to delete company", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ //nolint:wrapcheck
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent) //nolint:wrapcheck
}
