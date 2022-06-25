package operation

import (
	"database/sql"

	"github.com/Idea-Thrive/backend/internal/model"
	"go.uber.org/zap"
)

// Operation struct.
type Operation struct {
	DB     *sql.DB
	Logger *zap.Logger
}

// NewOperation function.
func NewOperation(db *sql.DB, logger *zap.Logger) *Operation {
	return &Operation{DB: db, Logger: logger}
}

func (o *Operation) UserGetByUsername(username string) (user model.User, err error) {
	return model.User{
		FirstName:   "first-name",
		LastName:    "last-name",
		Email:       "123@gmail.com",
		PhoneNumber: "0985858585",
		PhotoURL:    "",
		CompanyID:   "1234698755",
		PersonnelID: username,
		Gender:      "male",
		Role:        "employee",
	}, nil
}
