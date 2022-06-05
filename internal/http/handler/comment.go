package handler

import (
	"time"

	"github.com/Idea-Thrive/backend/internal/model"
	"github.com/Idea-Thrive/backend/internal/store"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Comment struct {
	Store  *store.Store
	Logger *zap.Logger
}

// Register function.
func (c Comment) Register(group fiber.Router) {
	group.Post("/", c.Create)
	group.Get("/", c.GetAll)
	group.Delete("/:id", c.Delete)
}

func (c Comment) Create(ctx *fiber.Ctx) error {
	comment := new(model.Comment)
	if err := ctx.BodyParser(comment); err != nil {
		c.Logger.Error("failed to parse comment", zap.Error(err))

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	if err := c.Store.CommentCreate(*comment); err != nil {
		c.Logger.Error("failed to create comment", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusCreated)
}
