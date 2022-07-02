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

	errRetrieve := u.DB.QueryRow("SELECT `category_id`, `title`, `description`, `is_approved`"+
		" `creator_id`, `company_id`, `created_at`, `updated_at` FROM `Idea` WHERE `id` = ?", id).Scan(
		&idea.CategoryID,
		&idea.Title,
		&idea.Description,
		&idea.IsApproved,
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
	queryString := "SELECT i.id, i.title, c.name, c.color, i.description, i.is_approved, " +
		"(SELECT AVG(cc.score) FROM Comment c INNER JOIN CriteriaComment cc ON c.id = cc.comment_id " +
		"WHERE c.idea_id = idea_id) AS score, i.category_id, i.creator_id, " +
		"i.created_at, i.updated_at FROM Idea i INNER JOIN Category c ON c.company_id = i.company_id WHERE 1"

	if companyID != "" {
		queryString += fmt.Sprintf(" AND i.company_id = %s", companyID)
	}

	queryString += fmt.Sprintf(" LIMIT %d OFFSET %d", size, offset)

	ideas, err := u.DB.Query(queryString)

	for ideas.Next() {
		var ideaItem model.Idea

		errScan := ideas.Scan(
			&ideaItem.ID,
			&ideaItem.Title,
			&ideaItem.CategoryName,
			&ideaItem.CategoryColor,
			&ideaItem.Description,
			&ideaItem.IsApproved,
			&ideaItem.Score,
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

func (u *Operation) IdeaEditStatus(id string) error {
	exec, err := u.DB.Exec("UPDATE `Idea` SET `is_approved` = true WHERE `id` = ?", id)

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
