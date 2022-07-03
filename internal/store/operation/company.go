package operation

import (
	"errors"
	"fmt"
	"time"

	"github.com/Idea-Thrive/backend/internal/model"
)

var errNotInsertedInCompanyTable = errors.New("not inserted in company table")

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
		return err
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
	errRetrieve := u.DB.QueryRow("SELECT `id`, `name`, `logo_url`, `owner_national_id`, "+
		"`owner_first_name`, `owner_last_name`, `created_at`, `updated_at` FROM `Company` WHERE `id` = ?", id).Scan(
		&company.ID,
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

// CompanyGetAll function.
func (u *Operation) CompanyGetAll(size, offset int) (res []model.Company, err error) {
	queryString := "SELECT `id`, `name`, `logo_url`, `owner_national_id`, `owner_first_name`, " +
		"`owner_last_name`, `created_at`, `updated_at` FROM `Company`"

	queryString += fmt.Sprintf(" LIMIT %d OFFSET %d", size, offset)

	companies, err := u.DB.Query(queryString)
	if err != nil {
		u.Logger.Error(err.Error())
	}

	for companies.Next() {
		var companyItem model.Company

		errScan := companies.Scan(
			&companyItem.ID,
			&companyItem.Name,
			&companyItem.LogoURL,
			&companyItem.OwnerNationalID,
			&companyItem.OwnerFirstName,
			&companyItem.OwnerLastName,
			&companyItem.CreatedAt,
			&companyItem.UpdatedAt,
		)

		if errScan != nil {
			return res, errScan
		}

		res = append(res, companyItem)
	}

	return res, nil
}

// CompanyUpdate function.
func (u *Operation) CompanyUpdate(companyID string, company model.Company) error {
	queryString := "UPDATE `Company` SET `updated_at` = ?, `name` = ?, `logo_url` = ?, `owner_national_id` = ?," +
		" `owner_first_name` = ?, `owner_last_name` = ? WHERE `id` = ?"

	res, err := u.DB.Exec(queryString,
		time.Now(),
		company.Name,
		company.LogoURL,
		company.OwnerNationalID,
		company.OwnerFirstName,
		company.OwnerLastName,
		companyID,
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
