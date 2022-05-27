package operation

import (
	"errors"
	"fmt"
	"time"

	"github.com/Idea-Thrive/backend/internal/model"
)

var (
	// errNotInsertedInUserTable error.
	errNotInsertedInUserTable = errors.New("not inserted in user table")
	// errUserAlreadyExistsInTable error.
	errUserAlreadyExistsInTable = errors.New("user already exists in user table")
	// errRetrieveDataFromUserTable error.
	errRetrieveDataFromUserTable = errors.New("cannot retrieve data from user table")
)

// UserCreate function.
func (u *Operation) UserCreate(user model.User) (err error) {

	currEmail := ""
	errRetrieve := u.DB.QueryRow("SELECT email from User WHERE email = ?", user.Email).Scan(&currEmail)

	if errRetrieve != nil {
		err = errRetrieveDataFromUserTable
	}

	if currEmail == user.Email {
		err = errUserAlreadyExistsInTable
	}

	queryString := "INSERT INTO User (first_name, last_name, email, phone_number, photo_url, company_id, personnel_id," +
		" gender, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := u.DB.Exec(queryString,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
		user.PhotoURL,
		user.CompanyID,
		user.PersonnelID,
		user.Gender,
		user.Role,
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
