package controller

import "github.com/gofiber/fiber/v2"

type LoanController interface {
	CreateLoan(c *fiber.Ctx) error
	ReturnLoan(c *fiber.Ctx) error
	ReadLoanDependencies(c *fiber.Ctx) error
}
