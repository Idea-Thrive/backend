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

func (u *Operation) CriteriaCreate(criteria model.Criteria) error {
	return nil
}

func (u *Operation) CriteriaGetAll(categoryID string) ([]model.Criteria, error) {
	c := model.Criteria{
		Name:       "c-1",
		CategoryID: "1",
	}

	criteria := make([]model.Criteria, 0)
	criteria = append(criteria, c)
	c.Name = "c-2"
	criteria = append(criteria, c)
	c.Name = "c-3"
	criteria = append(criteria, c)

	return criteria, nil
}

func (u *Operation) CriteriaDelete(id string) error {
	return nil
}

func (u *Operation) CategoryGet(id string) (*model.Category, error) {
	return &model.Category{
		Name:        "test-category",
		Description: "this is a test description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (u *Operation) CategoryGetAll(companyID string) ([]model.Category, error) {
	category := model.Category{
		Name:        "test-category",
		Description: "this is a test description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	categories := make([]model.Category, 0)
	categories = append(categories, category, category, category)

	return categories, nil
}

func (u *Operation) CategoryDelete(id string) error {
	return nil
}
