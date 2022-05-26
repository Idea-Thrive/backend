package operation

import (
	"database/sql"
	"go.uber.org/zap"
)

// Operation struct.
type Operation struct {
	DB     *sql.DB
	Logger *zap.Logger
}

// NewOperation function.
func NewOperation(DB *sql.DB, logger *zap.Logger) *Operation {
	return &Operation{DB: DB, Logger: logger}
}
