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
	queryString := "INSERT INTO `CategoryID` (`company_id`, `name`, `description`, `created_at`, `updated_at`)" +
		" VALUES (?, ?, ?, ?, ?)"

	result, err := u.DB.Exec(queryString,
		category.CompanyID,
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
	errRetrieve := u.DB.QueryRow("SELECT `company_id`, `name`, `description`, `created_at`, `updated_at` FROM `CategoryID` WHERE `id` = ?", id).Scan(
		&category.CompanyID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if errRetrieve != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (u *Operation) CategoryGetAll(companyID string) (res []model.Category, err error) {
	results, err := u.DB.Query("SELECT `name`, `description`, `created_at`, `updated_at` "+
		"FROM `CategoryID` WHERE `company_id` = ?", companyID)

	if err != nil {
		err = errRetrieveQueryError

		return res, err
	}

	for results.Next() {
		var categoryItem model.Category

		errScan := results.Scan(
			&categoryItem.Name,
			&categoryItem.Description,
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
	exec, err := u.DB.Exec("DELETE FROM `CategoryID` WHERE `id` = ?", id)

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
