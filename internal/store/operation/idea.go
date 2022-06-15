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
	fmt.Println("here")
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

func (u *Operation) IdeaGet(id string) (*model.Idea, error) {
	return &model.Idea{
		Title:         "good idea",
		Category:      1,
		Description:   "this is a very good idea",
		UpVoteCount:   10,
		DownVoteCount: -5,
		CreatorID:     "hassan",
		CompanyID:     "good-company",
		CreatedAt:     time.Now().String(),
		UpdatedAt:     time.Now().String(),
	}, nil
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
