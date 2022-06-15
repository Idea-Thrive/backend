package store

import "github.com/Idea-Thrive/backend/internal/model"

// Operation interface.
type Operation interface {
	Login(username, password string) (bool, error)

	UserCreate(user model.User) error
	UserGet(id string) (user model.User, err error)
	UserUpdate(id string, user model.User) error
	UserDelete(id string) error

	IdeaCreate(idea model.Idea) error
	IdeaGet(id string) (idea model.Idea, err error)
	IdeaGetAll(companyID, category string, size, offset int) ([]model.Idea, error)
	IdeaDelete(id string) error

	CompanyCreate(company model.Company) error
	CompanyGet(id string) (company model.Company, err error)
	CompanyUpdate(id string, company model.Company) error
	CompanyDelete(id string) error

	CategoryCreate(category model.Category) error
	CategoryGet(id string) (*model.Category, error)
	CategoryGetAll(companyID string) ([]model.Category, error)
	CategoryDelete(id string) error

	CriteriaCreate(criteria model.Criteria) error
	CriteriaGetAll(categoryID string) ([]model.Criteria, error)
	CriteriaDelete(id string) error

	CommentCreate(comment model.Comment) error
	CommentGetAll(ideaID string, scoreOnly bool, size, offset int) ([]model.Comment, error)
	CommentDelete(id string) error
}
