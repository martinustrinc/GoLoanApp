package table

import (
	"database/sql"
)

type AccountLoanStruct struct {
	Month                int     `json:"bulan"`
	PrincipalInstallment float64 `json:"angsuran_pokok"`
	InterestInstallment  float64 `json:"angsuran_bunga"`
	TotalInstallment     float64 `json:"total_angsuran"`
	RemainingInstallment float64 `json:"sisa_pinjaman"`
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
