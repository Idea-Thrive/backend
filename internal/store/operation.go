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
	IdeaGet(id string) (*model.Idea, error)
	IdeaGetAll(companyID, category string, size, offset int) ([]model.Idea, error)
	IdeaDelete(id string) error
}
