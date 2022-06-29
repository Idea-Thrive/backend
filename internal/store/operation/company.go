package operation

import (
	"errors"
	"github.com/Idea-Thrive/backend/internal/model"
	"time"
)

var (
	errNotInsertedInCompanyTable = errors.New("not inserted in company table")
)

// CompanyCreate function.
func (u *Operation) CompanyCreate(company model.Company) error {

	queryString := "INSERT INTO `Company`(`name`, `logo_url`, " +
		"`owner_national_id`, `owner_first_name`, `owner_last_name`, `created_at`, " +
		"`updated_at`)  VALUES (?, ?, ?, ?, ?, ?, ?)"

	result, err := u.DB.Exec(queryString,
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

// CompanyGet function.
func (u *Operation) CompanyGet(id string) (company model.Company, err error) {
	errRetrieve := u.DB.QueryRow("SELECT `name`, `logo_url`, `owner_national_id`, "+
		"`owner_first_name`, `owner_last_name`, `created_at`, `updated_at` FROM `Company` WHERE `id` = ?", id).Scan(
		&company.Name,
		&company.LogoURL,
		&company.OwnerNationalID,
		&company.OwnerFirstName,
		&company.OwnerLastName,
		&company.CreatedAt,
		&company.UpdatedAt,
	)

	if errRetrieve != nil {
		return model.Company{}, errNoRecordFound
	}

	return company, nil
}

// CompanyUpdate function.
func (u *Operation) CompanyUpdate(id string, company model.Company) error {
	queryString := "UPDATE `Company` SET `updated_at` = ?, `name` = ?, `logo_url` = ?, `owner_national_id` = ?," +
		" `owner_first_name` = ?, `owner_last_name` = ? WHERE `id` = ?"

	res, err := u.DB.Exec(queryString,
		time.Now(),
		company.Name,
		company.LogoURL,
		company.OwnerNationalID,
		company.OwnerFirstName,
		company.OwnerLastName,
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

// CompanyDelete function.
func (u *Operation) CompanyDelete(id string) error {
	exec, err := u.DB.Exec("DELETE FROM `Company` WHERE `id` = ?", id)

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
