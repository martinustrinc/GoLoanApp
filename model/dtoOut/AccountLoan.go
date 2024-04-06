package dtoOut

import (
	"GoLoanApp/model/table"
)

type LoanResponse struct {
	AccountCode string                    `json:"nomer_rekening"`
	AccountLoan []table.AccountLoanStruct `json:"angsuran"`
}
