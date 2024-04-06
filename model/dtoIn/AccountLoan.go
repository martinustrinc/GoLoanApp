package dtoIn

import (
	"GoLoanApp/model/table"
	"time"
)

type LoanRequest struct {
	CreditLimit  float64             `json:"plafond"`
	LoanDuration int                 `json:"lama_pinjaman"`
	InterestRate float64             `json:"suku_bunga"`
	StartDate    time.Time           `json:"tanggal_mulai"`
	Installment  []table.AccountLoan `json:"angsuran"`
}
