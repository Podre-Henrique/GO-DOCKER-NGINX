package controller

import (
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/controller/DTO/request"
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/model"
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/model/repo"
	"github.com/gofiber/fiber/v2"
)

func NewUserController(userRepo repo.UserRepo) UserController {
	return &userController{Repo: userRepo}
}

type userController struct {
	Repo repo.UserRepo
}

func (uc *userController) CreateUser(c *fiber.Ctx) error {
	var userRequest request.UserCreate
	if err := c.BodyParser(&userRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error")
	}

	user := &model.User{
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}
	newUser := uc.Repo.CreateUser(user, userRequest.Password)
	return c.JSON(newUser)
}

func (uc *userController) GetUser(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}

func (uc *userController) LoginUser(c *fiber.Ctx) error {
	panic("not implemented") // TODO: Implement
}
