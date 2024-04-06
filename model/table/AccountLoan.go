package table

import (
	"database/sql"
	"time"
)

type AccountLoanStruct struct {
	Month                int
	StartDate            time.Time
	PrincipalInstallment float64
	InterestInstallment  float64
	TotalInstallment     float64
	RemainingInstallment float64
	CreatedBy            int64
	UpdatedBy            int64
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type AccountLoan struct {
	ID                   sql.NullInt64
	Month                sql.NullInt64
	StartDate            sql.NullTime
	PrincipalInstallment sql.NullString
	InterestInstallment  sql.NullString
	TotalInstallment     sql.NullString
	RemainingInstallment sql.NullString
	CreatedBy            sql.NullInt64
	UpdatedBy            sql.NullInt64
	CreatedAt            sql.NullTime
	UpdatedAt            sql.NullTime
	Deleted              sql.NullBool
}
