package operation

import (
	"errors"
)

var (
	errNoRecordWithThisEmail = errors.New("there is no record with this email")
	errWrongPassword         = errors.New("wrong password entered")
)

// Login function.
func (u *Operation) Login(email, password string) (bool, error) {
	var realPassword string

	errRetrieve := u.DB.QueryRow("SELECT `password` FROM `User` WHERE `email` = ?", email).Scan(
		&realPassword,
	)

	if errRetrieve != nil {
		return false, errNoRecordWithThisEmail
	}

	if realPassword == password {
		return true, nil
	}

	return false, errWrongPassword
}
