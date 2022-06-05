package handler

import (
	"strconv"
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

func (c Comment) GetAll(ctx *fiber.Ctx) error {
	size, _ := strconv.Atoi(ctx.Query("size"))                         // optional
	offset, _ := strconv.Atoi(ctx.Query("offset"))                     // optional
	scoreOnly, _ := strconv.ParseBool(ctx.Query("scoreOnly", "false")) // optional
	ideaID := ctx.Query("idea_id")                                     // required

	if len(ideaID) == 0 {
		c.Logger.Error("idea_id is required")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "idea_id is required",
		})
	}

	comments, err := c.Store.CommentGetAll(ideaID, scoreOnly, size, offset)
	if err != nil {
		c.Logger.Error("failed to get comments", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(comments)
}

func (c Comment) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.Store.CommentDelete(id); err != nil {
		c.Logger.Error("failed to delete comment", zap.Error(err))

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
