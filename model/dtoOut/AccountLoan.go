package dtoOut

import (
	"GoLoanApp/model/table"
)

type LoanResponse struct {
	AccountCode string              `json:"kode_rekening"`
	AccountLoan []table.AccountLoan `json:"tabel_angsuran"`
}
