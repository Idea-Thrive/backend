package operation

import (
	"errors"
	"github.com/Idea-Thrive/backend/internal/model"
	"time"
)

var (
	errNotInsertedInCategoryTable = errors.New("not inserted in category table")
)

// CategoryCreate function.
func (u *Operation) CategoryCreate(category model.Category) (err error) {
	queryString := "INSERT INTO `Category` (`name`, `description`, `created_at`, `updated_at`)" +
		" VALUES (?, ?, ?, ?)"

	result, err := u.DB.Exec(queryString,
		category.Name,
		category.Description,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	lid, _ := result.LastInsertId()
	if lid == 0 {
		err = errNotInsertedInCategoryTable

		return err
	}

	return err
}
