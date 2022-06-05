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

}

func (c Company) Create(ctx *fiber.Ctx) error {
	company := new(model.Company)

	if err := ctx.BodyParser(company); err != nil {
		c.Logger.Error("failed to parse company", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	company.CreatedAt = time.Now()
	company.UpdatedAt = time.Now()

	if err := c.Store.CompanyCreate(*company); err != nil {
		c.Logger.Error("failed to create company", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}
