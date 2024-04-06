package repository

import (
	"GoLoanApp/model/table"
	"database/sql"
	"fmt"
)

type accountLoanDao struct {
	AbstractDAO
}

var AccountLoanDao = accountLoanDao{}.New()

func (input accountLoanDao) New() (output accountLoanDao) {
	output.FileName = "AccountLoanDao.go"
	output.TableName = "account_loan"
	return
}

func (input accountLoanDao) InsertAccountLoan(db *sql.Tx, userParam table.AccountLoan) (output int64, err error) {
	var (
		//funcName string
		query  string
		params []interface{}
	)

	query = fmt.Sprintf(`
		INSERT INTO %s 
		(
			id, month, total_installment, principal_installment, interest_installment, 
		 	remaining_installment, start_date, created_by, created_at, updated_by, updated_at, deleted
		) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		AccountLoanDao.TableName)

	params = append(params,
		userParam.ID.Int64, userParam.Month.Int64, userParam.StartDate.Time,
		userParam.PrincipalInstallment.String, userParam.InterestInstallment.String,
		userParam.TotalInstallment.String, userParam.RemainingInstallment.String,
		userParam.CreatedBy.Int64, userParam.CreatedAt.Time, userParam.UpdatedBy.Int64,
		userParam.UpdatedAt.Time, userParam.Deleted.Bool,
	)

	query += ` RETURNING id `
	results := db.QueryRow(query, params...)
	dbError := results.Scan(&output)
	if dbError != nil && dbError.Error() != sql.ErrNoRows.Error() {
		//err = GenerateInternalDBServerError(input.FileName, funcName, dbError)
		return
	}

	return
}
