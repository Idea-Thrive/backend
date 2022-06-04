package store

import "github.com/Idea-Thrive/backend/internal/model"

// Operation interface.
type Operation interface {
	Login(username, password string) (bool, error)

	UserCreate(user model.User) error

	UserGet(id string) (*model.User, error)

	UserDelete(id string) error
}
