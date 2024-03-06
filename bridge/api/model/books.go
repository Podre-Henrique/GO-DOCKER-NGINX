package model

import enums "github.com/Podre-Henrique/arquitetura-api/mvc/api/enum"

type Book struct {
	ISBN        uint64
	Name        string
	Description string
	Pages       uint16
	Quantity    uint16
	Category    enums.BookCategoryEnum
	Disponible  bool
}
