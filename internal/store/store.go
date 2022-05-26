package store

import (
	"github.com/Idea-Thrive/backend/internal/store/operation"
)

// Store struct.
type Store struct {
	Operation
}

// NewStore function.
func NewStore(connection *operation.Operation) *Store {
	return &Store{
		Operation: connection,
	}
}
