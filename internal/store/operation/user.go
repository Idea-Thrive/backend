package operation

import (
	"errors"
	"fmt"
	"github.com/Idea-Thrive/backend/internal/model"
	"time"
)

// errNotInsertedInUserTable error.
var errNotInsertedInUserTable = errors.New("not inserted in user table")

// UserCreate function.
func (u *Operation) UserCreate(user model.User) (err error) {
	queryString := "INSERT INTO User (first_name, last_name, email, phone_number, photo_url, personnel_id, gender, role, created_at, updated_at VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?))"

	result, err := u.DB.Exec(queryString,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
		user.PhotoURL,
		user.PersonnelID,
		user.Gender,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	lid, _ := result.LastInsertId()
	if lid == 0 {
		return errNotInsertedInUserTable
	}

	return fmt.Errorf("error: %w", err)
}
