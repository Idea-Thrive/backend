package operation

import (
	"errors"
	"github.com/Idea-Thrive/backend/internal/model"
	"time"
)

var (
	errNotInsertedInCommentTable = errors.New("not inserted in comment table")
)

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
		return err //nolint:wrapcheck
	}

	lid, _ := result.LastInsertId()
	if lid == 0 {
		err = errNotInsertedInCommentTable

		return err
	}

	return nil
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
