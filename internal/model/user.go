package model

import (
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

var (
	errNotInsertedInUserTable = errors.New("not inserted in user table")
)

type User struct {
	DB  *sql.DB
	Log *zap.SugaredLogger
}

type UserModel struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	PhotoURL    string `json:"photo_url"`
	PersonnelID string `json:"personnel_id"`
	Gender      string `json:"gender"`
	Role        string `json:"role"`
}

type UserImpl interface {
	Create(user UserModel) error
	//Update(id int, user UserModel) error
	//Delete(id int) error
}

func (u *UserModel) Create(user UserModel) (err error) {
	queryString := "INSERT INTO User (first_name, last_name, email, phone_number, photo_url, personnel_id, gender, role, created_at, updated_at VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?))"
	_ = queryString
	//result, err := u.DB.Exec(queryString,
	//	user.FirstName,
	//	user.LastName,
	//	user.Email,
	//	user.PhoneNumber,
	//	user.PhotoURL,
	//	user.PersonnelID,
	//	user.Gender,
	//	time.Now(),
	//	time.Now(),
	//)
	//
	//if err != nil {
	//	return err
	//}
	//
	//lid, _ := result.LastInsertId()
	//if lid == 0 {
	//	return errNotInsertedInUserTable
	//}

	return err
}
