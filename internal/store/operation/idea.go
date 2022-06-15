package operation

import (
	"errors"
	"time"

	"github.com/Idea-Thrive/backend/internal/model"
)

var (
	errNotInsertedInIdeaTable = errors.New("not inserted in idea table")
)

// IdeaCreate function.
func (u *Operation) IdeaCreate(idea model.Idea) (err error) {
	queryString := "INSERT INTO `Idea`(`category`, `title`, `description`," +
		" `up_vote`, `down_vote`, `creator_id`, `company_id`," +
		" `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := u.DB.Exec(queryString,
		idea.Category,
		idea.Title,
		idea.Description,
		idea.UpVoteCount,
		idea.DownVoteCount,
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

func (u *Operation) IdeaGet(id string) (idea model.Idea, err error) {

	errRetrieve := u.DB.QueryRow("SELECT `category`, `title`, `description`, `up_vote`, `down_vote`,"+
		" `creator_id`, `company_id`, `created_at`, `updated_at` FROM `Idea` WHERE `id` = ?", id).Scan(
		&idea.Category,
		&idea.Title,
		&idea.Description,
		&idea.UpVoteCount,
		&idea.DownVoteCount,
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

func (u *Operation) IdeaGetAll(companyID, category string, size, offset int) ([]model.Idea, error) {
	idea := model.Idea{
		Title:         "good idea",
		Category:      1,
		Description:   "this is a very good idea",
		UpVoteCount:   10,
		DownVoteCount: -5,
		CreatorID:     "hassan",
		CompanyID:     "good-company",
		CreatedAt:     time.Now().String(),
		UpdatedAt:     time.Now().String(),
	}

	ideas := make([]model.Idea, 0)
	ideas = append(ideas, idea, idea, idea)

	return ideas, nil
}

func (u *Operation) IdeaDelete(id string) error {
	return nil
}
