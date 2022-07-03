package operation

import (
	"errors"
	"fmt"
	"time"

	"github.com/Idea-Thrive/backend/internal/model"
)

var errNotInsertedInCommentTable = errors.New("not inserted in comment table")

// CommentCreate function.
func (u *Operation) CommentCreate(comment model.Comment) (err error) {
	queryString := "INSERT INTO `Comment`(`user_id`, `company_id`, `idea_id`, " +
		"`score`, `description`, `created_at`, " +
		"`updated_at`)  VALUES (?, ?, ?, ?, ?, ?, ?)"

	result, err := u.DB.Exec(queryString,
		comment.UserID,
		comment.CompanyID,
		comment.IdeaID,
		comment.Score,
		comment.Description,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	lid, _ := result.LastInsertId()
	if lid == 0 {
		err = errNotInsertedInCommentTable

		return err
	}

	return nil
}

// CommentGetAll function.
func (u *Operation) CommentGetAll(ideaID string, scoreOnly bool, size, offset int) (res []model.Comment, err error) {
	queryString := "SELECT  `id`, `company_id`, `user_id`, `score`, `description`," +
		" `created_at`, `updated_at` FROM `Comment` WHERE 1"

	if ideaID != "" {
		queryString += fmt.Sprintf(" AND idea_id = %s", ideaID)
	}

	if scoreOnly {
		queryString += fmt.Sprintf(" AND score <> %d", 0)
	}

	queryString += fmt.Sprintf(" LIMIT %d OFFSET %d", size, offset)

	ideas, err := u.DB.Query(queryString)
	if err != nil {
		u.Logger.Error(err.Error())
	}

	for ideas.Next() {
		var commentItem model.Comment

		errScan := ideas.Scan(
			&commentItem.ID,
			&commentItem.CompanyID,
			&commentItem.UserID,
			&commentItem.Score,
			&commentItem.Description,
			&commentItem.CreatedAt,
			&commentItem.UpdatedAt,
		)

		if errScan != nil {
			return res, errScan
		}

		res = append(res, commentItem)
	}

	return res, nil
}

// CommentDelete function.
func (u *Operation) CommentDelete(id string) error {
	exec, err := u.DB.Exec("DELETE FROM `Comment` WHERE `id` = ?", id)
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
