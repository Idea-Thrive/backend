package store

import (
	"github.com/Idea-Thrive/backend/internal/mysql/operation"
)

type Operations interface {
	Login(username, password string) (bool, error)
}

type Store struct {
	Operations
	DB *operation.Operation
}

func NewStore(connection *operation.Operation) *Store {
	return &Store{
		DB: connection,
	}
}
