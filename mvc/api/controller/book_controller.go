package controller

type BookController interface {
	CreateBook()
	AddBooks()
	ReadDisponibleBooks()
	ReadLoanUserBooks()
}
