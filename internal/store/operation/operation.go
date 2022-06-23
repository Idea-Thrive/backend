package operation

import (
	"database/sql"
	"time"

	"github.com/Idea-Thrive/backend/internal/model"
	"go.uber.org/zap"
)

// Operation struct.
type Operation struct {
	DB     *sql.DB
	Logger *zap.Logger
}

// NewOperation function.
func NewOperation(db *sql.DB, logger *zap.Logger) *Operation {
	return &Operation{DB: db, Logger: logger}
}

func (u *Operation) CommentCreate(comment model.Comment) error {
	return nil
}

func (u *Operation) CommentGetAll(ideaID string, scoreOnly bool, size, offset int) ([]model.Comment, error) {
	comment := model.Comment{
		Score:       2,
		Description: "yes",
		UserID:      123123,
		CompanyID:   2352354,
		IdeaID:      423534543,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return []model.Comment{comment, comment, comment}, nil
}

func (u *Operation) CommentDelete(id string) error {
	return nil
}
