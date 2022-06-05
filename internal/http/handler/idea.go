package handler

import (
	"strconv"
	"time"

	"github.com/Idea-Thrive/backend/internal/model"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
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
	group.Delete("/:id", i.Delete)
}

func (i Idea) Create(ctx *fiber.Ctx) error {
	idea := new(model.Idea)

	if err := ctx.BodyParser(&idea); err != nil {
		i.Logger.Error("failed to parse body", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse body",
		}) //nolint:wrapcheck
	}

	idea.CreatedAt = time.Now()
	idea.UpdatedAt = time.Now()

	if err := i.Store.IdeaCreate(*idea); err != nil {
		i.Logger.Error("failed to create idea", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create idea",
		}) //nolint:wrapcheck
	}

	return ctx.SendStatus(fiber.StatusOK)
}

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

func (i Idea) GetAll(ctx *fiber.Ctx) error {
	size, _ := strconv.Atoi(ctx.Query("size", "100"))
	offset, _ := strconv.Atoi(ctx.Query("offset", "0"))

	ideas, err := i.Store.IdeaGetAll(size, offset)
	if err != nil {
		i.Logger.Error("failed to get ideas", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get ideas",
		}) //nolint:wrapcheck
	}

	return ctx.Status(fiber.StatusOK).JSON(ideas) //nolint:wrapcheck
}

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
