package store

import (
	"github.com/Idea-Thrive/backend/internal/store/operation"
)

type Store struct {
	Operation
}

func NewStore(connection *operation.Operation) *Store {
	return &Store{
		Operation: connection,
	}
}
