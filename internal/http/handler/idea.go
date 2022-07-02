package handler

import (
	"github.com/Idea-Thrive/backend/internal/http/request"
	"github.com/Idea-Thrive/backend/internal/model"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
)

// Idea struct.
type Idea struct {
	Store  *store.Store
	Logger *zap.Logger
}

// Register function.
func (i Idea) Register(group fiber.Router) {
	group.Post("/", i.Create)
	group.Get("/:id", i.Get)
	group.Get("/", i.GetAll)
	group.Put("/status/:id", i.EditStatus)
	group.Delete("/:id", i.Delete)
}

// Create function.
func (i Idea) Create(ctx *fiber.Ctx) error {
	req := new(request.IdeaCreation)

	if err := ctx.BodyParser(&req); err != nil {
		i.Logger.Error("failed to parse body", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse body",
		}) //nolint:wrapcheck
	}

	idea := model.Idea{
		CategoryID:  req.CategoryID,
		Title:       req.Title,
		Description: req.Description,
		CreatorID:   req.CreatorID,
		CompanyID:   req.CompanyID,
	}

	if err := i.Store.IdeaCreate(idea); err != nil {
		i.Logger.Error("failed to create idea", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create idea",
		}) //nolint:wrapcheck
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// Get function,
func (i Idea) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	idea, err := i.Store.IdeaGet(id)
	if err != nil {
		i.Logger.Error("failed to get idea", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get idea",
		}) //nolint:wrapcheck
	}

	return ctx.Status(fiber.StatusOK).JSON(idea) //nolint:wrapcheck
}

// GetAll function.
func (i Idea) GetAll(ctx *fiber.Ctx) error {
	size, _ := strconv.Atoi(ctx.Query("size", "100"))   // optional
	offset, _ := strconv.Atoi(ctx.Query("offset", "0")) // optional
	companyID := ctx.Query("company_id")                // required

	if len(companyID) == 0 {
		i.Logger.Warn("company_id cannot be null")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "company_id cannot be null",
		}) //nolint:wrapcheck
	}

	ideas, err := i.Store.IdeaGetAll(companyID, size, offset)
	if err != nil {
		i.Logger.Error("failed to get ideas", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get ideas",
		}) //nolint:wrapcheck
	}

	return ctx.Status(fiber.StatusOK).JSON(ideas) //nolint:wrapcheck
}

// EditStatus function.
func (i Idea) EditStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := i.Store.IdeaEditStatus(id); err != nil {
		i.Logger.Error("failed to change status idea", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to change status of idea",
		}) //nolint:wrapcheck
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// Delete function.
func (i Idea) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := i.Store.IdeaDelete(id); err != nil {
		i.Logger.Error("failed to delete idea", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete idea",
		}) //nolint:wrapcheck
	}

	return ctx.SendStatus(fiber.StatusNoContent) //nolint:wrapcheck
}
