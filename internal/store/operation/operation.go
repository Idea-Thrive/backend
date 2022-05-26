package operation

import (
	"database/sql"
	"go.uber.org/zap"
)

type Operation struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func (u *Operation) Login(username, password string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewOperation(DB *sql.DB, logger *zap.Logger) *Operation {
	return &Operation{DB: DB, Logger: logger}
}
