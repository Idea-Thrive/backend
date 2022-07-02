package store

import "github.com/Idea-Thrive/backend/internal/model"

// Operation interface.
type Operation interface {
	Login(username, password string) (bool, error)

	UserCreate(user model.User) error
	UserGet(id string) (user model.User, err error)
	UserGetAll(size, offset int) (users []model.User, err error)
	UserGetByUsername(username string) (user model.User, err error)
	UserUpdate(id string, user model.User) error
	UserChangeRole(id string, newUserRole string) error
	UserDelete(id string) error

	IdeaCreate(idea model.Idea) error
	IdeaGet(id string) (idea model.Idea, err error)
	IdeaGetAll(companyID string, size, offset int) ([]model.Idea, error)
	IdeaEditStatus(id string) error
	IdeaDelete(id string) error

	CompanyCreate(company model.Company) error
	CompanyGet(id string) (company model.Company, err error)
	CompanyGetAll(size, offset int) ([]model.Company, error)
	CompanyUpdate(id string, company model.Company) error
	CompanyDelete(id string) error

	CategoryCreate(category model.Category) error
	CategoryGet(id string) (model.Category, error)
	CategoryGetAll(companyID string) ([]model.Category, error)
	CategoryDelete(id string) error

	CriteriaCreate(criteria model.Criteria) error
	CriteriaGetAll(categoryID string) ([]model.Criteria, error)
	CriteriaDelete(id string) error

	CommentCreate(comment model.Comment) error
	CommentGetAll(ideaID string, scoreOnly bool, size, offset int) ([]model.Comment, error)
	CommentDelete(id string) error
}
