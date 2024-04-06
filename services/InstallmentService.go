package services

import (
	"GoLoanApp/model/dtoIn"
	"GoLoanApp/model/table"
)

func CalculateInstallment(request dtoIn.LoanRequest) (tableInstallment []table.AccountLoanStruct, err error) {
	//installment := make([]dtoIn.AccountLoan, request.LoanDuration)
	//date := time.Time{}

	// Calculate PrincipalInstallment
	monthlyPrincipal := request.CreditLimit / float64(request.LoanDuration)
	// Calculate InterestInstallment
	monthlyInterest := request.InterestRate / 12
	// Calculate AccountLoan
	//installment := request.CreditLimit * (monthlyInterest * math.Pow(1+monthlyInterest, float64(request.LoanDuration))) / (math.Pow(1+monthlyInterest, float64(request.LoanDuration)) - 1)

	for i := 0; i < request.LoanDuration; i++ {
		// Calculate Remaining Loan
		remainingInstallment := request.CreditLimit - float64(i-1)*monthlyPrincipal

		// Calculate Interest Rate
		interestRate := remainingInstallment * monthlyInterest

		// Calculate total AccountLoan
		totalInstallment := monthlyPrincipal + interestRate

		// Add data ke table Account
		tableInstallment = append(tableInstallment, table.AccountLoanStruct{
			Month:                i,
			PrincipalInstallment: monthlyPrincipal,
			InterestInstallment:  interestRate,
			TotalInstallment:     totalInstallment,
			RemainingInstallment: remainingInstallment,
		})
	}

	return tableInstallment, err
}

func ValidateRequestCreateLoanAccount(request dtoIn.LoanRequest) string {
	if request.CreditLimit <= 0 || request.LoanDuration <= 0 || request.InterestRate <= 0 {
		return "Invalid Input Data, Data must higher than 0!"
	}
	return ""
}
