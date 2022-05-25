package store

import "github.com/Idea-Thrive/backend/internal/mysql"

type Operations interface {
	Login(username, password string) (bool, error)
}

type Store struct {
	Operations
	Connection mysql.Mysql
}

func NewStore(connection *mysql.Mysql) *Store {
	return &Store{}
}
