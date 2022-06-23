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

// CategoryGet function.
func (u *Operation) CategoryGet(id string) (category model.Category, err error) {
	errRetrieve := u.DB.QueryRow("SELECT `name`, `description`, `created_at`, `updated_at` FROM `Category` WHERE `id` = ?", id).Scan(
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if errRetrieve != nil {
		return model.Category{}, nil
	}

	return category, nil
}

// CategoryDelete function.
func (u *Operation) CategoryDelete(id string) error {
	exec, err := u.DB.Exec("DELETE FROM `Category` WHERE `id` = ?", id)

	if err != nil {
		return err
	}

	rAffected, err := exec.RowsAffected()
	if err != nil {
		err = errCallingRowsAffected

		return err
	}

	if rAffected == 0 {
		err = errNoRowsAffected

		return err
	}
	return nil
}
