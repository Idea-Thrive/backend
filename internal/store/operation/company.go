package operation

import (
	"errors"
	"github.com/Idea-Thrive/backend/internal/model"
	"time"
)

var (
	errRepetitiveCompanyID       = errors.New("company with this company id already exists")
	errNotInsertedInCompanyTable = errors.New("not inserted in company table")
)

func (u *Operation) CompanyCreate(company model.Company) error {
	currCompanyID := ""
	errRetrieve := u.DB.QueryRow("SELECT `company_id` FROM `Company` WHERE `company_id` = ?",
		company.CompanyID).Scan(&currCompanyID)

	if errRetrieve != nil {
		u.Logger.Error("no company id found with this id")
	}

	if currCompanyID == company.CompanyID {
		return errRepetitiveCompanyID
	}

	queryString := "INSERT INTO `Company`(`company_id`, `name`, `logo_url`, " +
		"`owner_national_id`, `owner_first_name`, `owner_last_name`, `created_at`, " +
		"`updated_at`)  VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := u.DB.Exec(queryString,
		company.CompanyID,
		company.Name,
		company.LogoURL,
		company.OwnerNationalID,
		company.OwnerFirstName,
		company.OwnerLastName,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err //nolint:wrapcheck
	}

	lid, _ := result.LastInsertId()
	if lid == 0 {
		err = errNotInsertedInCompanyTable

		return err
	}

	return nil
}