package operation

import (
	"time"

	"github.com/Idea-Thrive/backend/internal/model"
)

func (u *Operation) IdeaCreate(idea model.Idea) error {
	return nil
}

func (u *Operation) IdeaGet(id string) (*model.Idea, error) {
	return &model.Idea{
		Title:         "good idea",
		Category:      "others",
		Description:   "this is a very good idea",
		UpVoteCount:   10,
		DownVoteCount: -5,
		CreatorID:     "hassan",
		CompanyID:     "good-company",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}

func (u *Operation) IdeaGetAll(companyID, category string, size, offset int) ([]model.Idea, error) {
	idea := model.Idea{
		Title:         "good idea",
		Category:      "others",
		Description:   "this is a very good idea",
		UpVoteCount:   10,
		DownVoteCount: -5,
		CreatorID:     "hassan",
		CompanyID:     "good-company",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	ideas := make([]model.Idea, 0)
	ideas = append(ideas, idea, idea, idea)

	return ideas, nil
}

func (u *Operation) IdeaDelete(id string) error {
	return nil
}
