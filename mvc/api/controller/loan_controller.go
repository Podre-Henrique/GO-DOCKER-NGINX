package controller

type LoanController interface {
	CreateLoan()
	ReturnLoan()
	ReadLoanDependencies()
}
