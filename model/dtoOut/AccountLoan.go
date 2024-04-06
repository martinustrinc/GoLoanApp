package dtoOut

import (
	"GoLoanApp/model/table"
)

type LoanResponse struct {
	AccountCode string                    `json:"kode_rekening"`
	AccountLoan []table.AccountLoanStruct `json:"tabel_angsuran"`
}
