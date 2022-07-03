package handler

import (
	"strconv"
	"time"

	"github.com/Idea-Thrive/backend/internal/http/request"
	"github.com/Idea-Thrive/backend/internal/model"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Company function.
type Company struct {
	Store  *store.Store
	Logger *zap.Logger
}

// Register function.
func (c Company) Register(group fiber.Router) {
	group.Post("/", c.Create)
	group.Get("/:id", c.Get)
	group.Get("/", c.GetAll)
	group.Delete("/:id", c.Delete)
	group.Put("/:id", c.Update)
}

// Create function.
func (c Company) Create(ctx *fiber.Ctx) error {
	req := new(request.CompanyCreation)

	if err := ctx.BodyParser(req); err != nil {
		c.Logger.Error("failed to parse req", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	company := model.Company{
		Name:            req.Name,
		LogoURL:         req.LogoURL,
		OwnerNationalID: req.OwnerNationalID,
		OwnerFirstName:  req.OwnerFirstName,
		OwnerLastName:   req.OwnerLastName,
		CreatedAt:       time.Now().String(),
		UpdatedAt:       time.Now().String(),
	}

	if err := c.Store.CompanyCreate(company); err != nil {
		c.Logger.Error("failed to create company", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// Get function.
func (c Company) Get(ctx *fiber.Ctx) error {
	id := ctx.AllParams()["id"]

	company, err := c.Store.CompanyGet(id)
	if err != nil {
		c.Logger.Error("failed to get company", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(company)
}

// Update function.
func (c Company) Update(ctx *fiber.Ctx) error {
	companyID := ctx.AllParams()["companyID"]

	req := new(request.CompanyCreation)

	if err := ctx.BodyParser(req); err != nil {
		c.Logger.Error("failed to parse req", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	company := model.Company{
		Name:            req.Name,
		LogoURL:         req.LogoURL,
		OwnerNationalID: req.OwnerNationalID,
		OwnerFirstName:  req.OwnerFirstName,
		OwnerLastName:   req.OwnerLastName,
		CreatedAt:       time.Now().String(),
		UpdatedAt:       time.Now().String(),
	}

	if err := c.Store.CompanyUpdate(companyID, company); err != nil {
		c.Logger.Error("failed to update company", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// Delete function.
func (c Company) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.Store.CompanyDelete(id); err != nil {
		c.Logger.Error("failed to delete company", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// GetAll function.
func (c Company) GetAll(ctx *fiber.Ctx) error {
	size, _ := strconv.Atoi(ctx.Query("size", "100"))   // optional
	offset, _ := strconv.Atoi(ctx.Query("offset", "0")) // optional

	companies, err := c.Store.CompanyGetAll(size, offset)
	if err != nil {
		c.Logger.Error("failed to get companies", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get companies",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(companies)
}
