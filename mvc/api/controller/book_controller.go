package controller

import (
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/model/repo"
	"github.com/gofiber/fiber/v2"
)

type BookController interface {
	CreateBook(c *fiber.Ctx) error
	AddBooks(c *fiber.Ctx) error
	ReadDisponibleBooks(c *fiber.Ctx) error
	ReadLoanUserBooks(c *fiber.Ctx) error
}

type bookController struct {
	BookRepo *repo.BookRepo
}

func NewBookRepo(bookRepo *repo.BookRepo) BookController {
	return &bookController{BookRepo: bookRepo}
}
