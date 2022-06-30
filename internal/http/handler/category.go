package handler

import (
	"github.com/Idea-Thrive/backend/internal/http/request"
	"time"

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
	group.Get("/:id", c.Get)
	group.Get("/", c.GetAll)
	group.Delete("/:id", c.Delete)
}

func (c Category) Create(ctx *fiber.Ctx) error {
	req := new(request.CategoryCreation)

	if err := ctx.BodyParser(&req); err != nil {
		c.Logger.Error("failed to parse body", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	category := model.Category{
		CompanyID:   req.CompanyID,
		Name:        req.Name,
		Description: req.Color,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	if err := c.Store.CategoryCreate(category); err != nil {
		c.Logger.Error("failed to create category", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusCreated) //nolint:wrapcheck
}

func (c Category) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	category, err := c.Store.CategoryGet(id)
	if err != nil {
		c.Logger.Error("failed to get category", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(category) //nolint:wrapcheck
}

func (c Category) GetAll(ctx *fiber.Ctx) error {
	companyID := ctx.Query("company_id")

	if len(companyID) == 0 {
		c.Logger.Error("company_id is required")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{ //nolint:wrapcheck
			"error": "company_id is required",
		})
	}

	categories, err := c.Store.CategoryGetAll(companyID)
	if err != nil {
		c.Logger.Error("failed to get categories", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ //nolint:wrapcheck
			"error": err.Error(),
		})
	}

	return ctx.JSON(categories) //nolint:wrapcheck
}

func (c Category) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.Store.CategoryDelete(id); err != nil {
		c.Logger.Error("failed to delete category", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ //nolint:wrapcheck
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent) //nolint:wrapcheck
}
