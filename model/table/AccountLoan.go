package table

import "time"

type AccountLoan struct {
	Month                int       `json:"bulan"`
	StartDate            time.Time `json:"tanggal_mulai"`
	PrincipalInstallment float64   `json:"angsuran_pokok"`
	InterestInstallment  float64   `json:"angsuran_bunga"`
	TotalInstallment     float64   `json:"total_angsuran"`
	RemainingInstallment float64   `json:"sisa_angsuran"`
	CreatedBy            int64     `json:"created_by"`
	UpdatedBy            int64     `json:"updated_by"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
