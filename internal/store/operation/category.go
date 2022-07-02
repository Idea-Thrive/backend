package operation

import (
	"errors"
	"github.com/Idea-Thrive/backend/internal/model"
	"time"
)

var (
	errNotInsertedInCategoryTable = errors.New("not inserted in category table")
	errRetrieveQueryError         = errors.New("error in retrieve query error")
)

// CategoryCreate function.
func (u *Operation) CategoryCreate(category model.Category) (err error) {
	queryString := "INSERT INTO `Category` (`company_id`, `name`, `color`, `created_at`, `updated_at`)" +
		" VALUES (?, ?, ?, ?, ?)"

	result, err := u.DB.Exec(queryString,
		category.CompanyID,
		category.Name,
		category.Color,
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
	errRetrieve := u.DB.QueryRow("SELECT `company_id`, `name`, `color`, `created_at`, `updated_at` FROM `Category` WHERE `id` = ?", id).Scan(
		&category.CompanyID,
		&category.Name,
		&category.Color,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if errRetrieve != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (u *Operation) CategoryGetAll(companyID string) (res []model.Category, err error) {
	results, err := u.DB.Query("SELECT `id`, `name`, `color`, `created_at`, `updated_at` "+
		"FROM `Category` WHERE `company_id` = ?", companyID)

	if err != nil {
		err = errRetrieveQueryError

		return res, err
	}

	for results.Next() {
		var categoryItem model.Category

		errScan := results.Scan(
			&categoryItem.ID,
			&categoryItem.Name,
			&categoryItem.Color,
			&categoryItem.CreatedAt,
			&categoryItem.UpdatedAt,
		)
		if errScan != nil {
			return res, errScan
		}

		res = append(res, categoryItem)
	}

	return res, err
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
