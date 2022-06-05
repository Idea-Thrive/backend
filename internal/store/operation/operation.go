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

func (u *Operation) CategoryCreate(category model.Category) error {
	return nil
}

func (u *Operation) CategoryGet(id string) (*model.Category, error) {
	return &model.Category{
		Name:        "test-category",
		Description: "this is a test description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (u *Operation) CategoryGetAll(companyID, id string) ([]model.Category, error) {
	category := model.Category{
		Name:        "test-category",
		Description: "this is a test description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	categories := make([]model.Category, 0)
	categories = append(categories, category, category, category)

	return categories, nil
}

func (u *Operation) CategoryDelete(id string) error {
	return nil
}

func (u *Operation) CompanyCreate(company model.Company) error {
	return nil
}

func (u *Operation) CompanyGet(id string) (*model.Company, error) {
	return &model.Company{
		Name:            "test company",
		LogoURL:         "url",
		OwnerNationalID: "2520202020202",
		OwnerFirstName:  "test-first",
		OwnerLastName:   "test-last",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, nil
}

func (u *Operation) CompanyDelete(id string) error {
	return nil
}
