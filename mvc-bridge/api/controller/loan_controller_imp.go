package controller

import (
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/model/repo"
	"github.com/gofiber/fiber/v2"
)

type loanController struct {
	Repo repo.LoanRepo
}

func (lc *loanController) CreateLoan(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (lc *loanController) ReturnLoan(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (lc *loanController) ReadLoanDependencies(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
