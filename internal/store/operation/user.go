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
	// errNoRecordFound error.
	errNoRecordFound = errors.New("no record found for this user")
	// errNoRowsAffected error.
	errNoRowsAffected = errors.New("no rows affected")
	// errNoRowsUpdated error.
	errNoRowsUpdated = errors.New("no rows updated")
	// errCallingRowsAffected error.
	errCallingRowsAffected = errors.New("error in calling rowsAffected function")
)

// UserCreate function.
func (u *Operation) UserCreate(user model.User) (err error) {

	currEmail := ""
	errRetrieve := u.DB.QueryRow("SELECT email from User WHERE email = ?", user.Email).Scan(&currEmail)

	if currEmail == user.Email {
		err = errUserAlreadyExistsInTable

		return err
	}

	if errRetrieve != nil {
		u.Logger.Error("user with this email doesn't exist")
	}

	queryString := "INSERT INTO `User` (`first_name`, `last_name`, `email`, `password`, `phone_number`, `photo_url`, `company_id`, `personnel_id`," +
		" `gender`, `role`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := u.DB.Exec(queryString,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
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
		return err
	}

	lid, _ := result.LastInsertId()
	if lid == 0 {
		err = errNotInsertedInUserTable

		return err
	}

	return err
}

// UserGet function.
func (u *Operation) UserGet(id string) (user model.User, err error) {
	errRetrieve := u.DB.QueryRow("SELECT `id`, `first_name`, `last_name`,"+
		" `email`, `phone_number`, `photo_url`, `company_id`, `personnel_id`, `gender`, `role` FROM `User` WHERE `id` = ?", id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PhoneNumber,
		&user.PhotoURL,
		&user.CompanyID,
		&user.PersonnelID,
		&user.Gender,
		&user.Role,
	)

	if errRetrieve != nil {
		return model.User{}, errNoRecordFound
	}

	return user, nil
}

// UserGetAll function.
func (u *Operation) UserGetAll(size, offset int) (res []model.User, err error) {
	queryString := "SELECT `id`, `first_name`, `last_name`," +
		" `email`, `phone_number`, `photo_url`, `company_id`, `personnel_id`, `gender`, `role` FROM `User` "

	queryString += fmt.Sprintf(" LIMIT %d OFFSET %d", size, offset)

	users, err := u.DB.Query(queryString)
	if err != nil {
		u.Logger.Error(err.Error())
	}

	for users.Next() {
		var user model.User

		errScan := users.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.PhoneNumber,
			&user.PhotoURL,
			&user.CompanyID,
			&user.PersonnelID,
			&user.Gender,
			&user.Role,
		)

		if errScan != nil {
			return res, errScan
		}

		res = append(res, user)
	}

	return res, nil
}

// UserGetByUsername function.
func (u *Operation) UserGetByUsername(username string) (user model.User, err error) {
	errRetrieve := u.DB.QueryRow("SELECT `id`, `first_name`, `last_name`, `email`, `phone_number`, `photo_url`,"+
		" `company_id`, `personnel_id`, `gender`, `role` FROM `User` WHERE `email` = ?", username).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PhoneNumber,
		&user.PhotoURL,
		&user.CompanyID,
		&user.PersonnelID,
		&user.Gender,
		&user.Role,
	)
	if errRetrieve != nil {
		return user, errRetrieve
	}

	return user, nil
}

// UserUpdate function.
func (u *Operation) UserUpdate(id string, user model.User) error {

	queryString := "UPDATE `User` SET `updated_at` = ?, `first_name` = ?, `last_name` = ?, `email` = ?, `password` = ?, `phone_number` = ?," +
		" `photo_url` = ?, `company_id` = ?, `personnel_id` = ?, `gender` = ?, `role` = ? WHERE `id` = ?"

	res, err := u.DB.Exec(queryString,
		time.Now(),
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.PhoneNumber,
		user.PhotoURL,
		user.CompanyID,
		user.PersonnelID,
		user.Gender,
		user.Role,
		id,
	)
	if err != nil {
		err = errNoRowsUpdated

		return err
	}

	rAffected, err := res.RowsAffected()
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

// UserDelete function.
func (u *Operation) UserDelete(id string) error {
	exec, err := u.DB.Exec("DELETE FROM `User` WHERE `id` = ?", id)

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
