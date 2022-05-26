package store

import (
	"github.com/Idea-Thrive/backend/internal/mysql/operation"
)

type Store struct {
	Operation
	DB *operation.Operation
}

func NewStore(connection *operation.Operation) *Store {
	return &Store{
		DB: connection,
	}
}
