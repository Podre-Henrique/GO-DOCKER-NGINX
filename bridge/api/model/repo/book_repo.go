package repo

import (
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/model"
)

type BookRepo interface {
	CreateBook(b *model.Book)
	AddBooks(isbn uint64, quantity uint16)
	ReadDisponibleBooks() *[]model.Book
	ReadLoanUserBooks(userId uint64) *[]model.Book
}
