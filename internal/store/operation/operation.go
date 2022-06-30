package operation

import (
	"database/sql"
	"time"

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

func (o *Operation) CompanyGetAll(size, offset int) ([]model.Company, error) {
	company := model.Company{
		ID:              "123123123",
		Name:            "c-1",
		LogoURL:         "",
		OwnerNationalID: "2522222222",
		OwnerFirstName:  "jafar",
		OwnerLastName:   "tehrani",
		CreatedAt:       time.Now().String(),
		UpdatedAt:       time.Now().String(),
	}

	return []model.Company{company, company, company, company}, nil
}
