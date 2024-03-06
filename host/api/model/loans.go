package model

type Loan struct {
	Id          uint32
	UserId      uint64
	BookISBN    uint64
	LoanDate    string
	DueDate     string
	WasReturned bool
}
