package handler

import (
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

}
