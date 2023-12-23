package repo

import "github.com/Podre-Henrique/arquitetura-api/mvc/api/model"

type LoanRepo interface {
	CreateLoan(b *model.Loan)
	ReturnLoan(userId uint64, isbn uint64)
	ReadLoanDependencies() *[]model.Book
}
