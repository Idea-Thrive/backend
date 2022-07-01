package operation

import (
	"errors"
	"fmt"
	"time"

	"github.com/Idea-Thrive/backend/internal/model"
)

var (
	errNotInsertedInIdeaTable = errors.New("not inserted in idea table")
)

// IdeaCreate function.
func (u *Operation) IdeaCreate(idea model.Idea) (err error) {
	queryString := "INSERT INTO `Idea`(`category_id`, `title`, `description`," +
		" `creator_id`, `company_id`," +
		" `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?, ?)"

	result, err := u.DB.Exec(queryString,
		idea.CategoryID,
		idea.Title,
		idea.Description,
		idea.CreatorID,
		idea.CompanyID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	lid, _ := result.LastInsertId()
	if lid == 0 {
		err = errNotInsertedInIdeaTable

		return err
	}

	return err
}

// IdeaGet function.
func (u *Operation) IdeaGet(id string) (idea model.Idea, err error) {

	errRetrieve := u.DB.QueryRow("SELECT `category_id`, `title`, `description`,"+
		" `creator_id`, `company_id`, `created_at`, `updated_at` FROM `Idea` WHERE `id` = ?", id).Scan(
		&idea.CategoryID,
		&idea.Title,
		&idea.Description,
		&idea.CreatorID,
		&idea.CompanyID,
		&idea.CreatedAt,
		&idea.UpdatedAt,
	)

	if errRetrieve != nil {
		return model.Idea{}, errNoRecordFound
	}

	return idea, nil

}

// IdeaGetAll function.
func (u *Operation) IdeaGetAll(companyID string, size, offset int) (res []model.Idea, err error) {
	queryString := "SELECT `id`, `title`, `description`, `category_id`, `creator_id`, " +
		"`created_at`, `updated_at` FROM `Idea` WHERE 1"

	if companyID != "" {
		queryString += fmt.Sprintf(" AND `company_id` = %s", companyID)
	}

	queryString += fmt.Sprintf(" LIMIT %d OFFSET %d", size, offset)

	ideas, err := u.DB.Query(queryString)

	for ideas.Next() {
		var ideaItem model.Idea

		errScan := ideas.Scan(
			&ideaItem.ID,
			&ideaItem.Title,
			&ideaItem.Description,
			&ideaItem.CategoryID,
			&ideaItem.CreatorID,
			&ideaItem.CreatedAt,
			&ideaItem.UpdatedAt,
		)

		if errScan != nil {
			return res, errScan
		}

		res = append(res, ideaItem)
	}

	return res, nil
}

// IdeaDelete function.
func (u *Operation) IdeaDelete(id string) error {
	exec, err := u.DB.Exec("DELETE FROM `Idea` WHERE `id` = ?", id)

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
