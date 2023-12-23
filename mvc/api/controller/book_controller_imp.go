package controller

import (
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/model/repo"
	"github.com/gofiber/fiber/v2"
)

type bookController struct {
	BookRepo repo.BookRepo
}

func (bc *bookController) CreateBook(c *fiber.Ctx) {
}

func (bc *bookController) AddBooks(c *fiber.Ctx) {
}

func (bc *bookController) ReadDisponibleBooks(c *fiber.Ctx) {
}

func (bc *bookController) ReadLoanUserBooks(c *fiber.Ctx) {
}
