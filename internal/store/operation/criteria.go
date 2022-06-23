package operation

import (
	"errors"
	"github.com/Idea-Thrive/backend/internal/model"
)

var (
	errNotInsertedInCriteriaTable = errors.New("not inserted in criteria table")
)

// CriteriaCreate function.
func (u *Operation) CriteriaCreate(criteria model.Criteria) error {
	queryString := "INSERT INTO `Criteria` (`category_id`, `name`) VALUES (?, ?)"

	result, err := u.DB.Exec(queryString,
		criteria.CategoryID,
		criteria.Name,
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	lid, _ := result.LastInsertId()
	if lid == 0 {
		err = errNotInsertedInCriteriaTable

		return err
	}

	return err
}

// CriteriaDelete function.
func (u *Operation) CriteriaDelete(id string) error {
	exec, err := u.DB.Exec("DELETE FROM `Criteria` WHERE `id` = ?", id)

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
