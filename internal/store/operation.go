package store

import "github.com/Idea-Thrive/backend/internal/model"

type Operation interface {
	Login(username, password string) (bool, error)

	UserCreate(user model.User) error
}