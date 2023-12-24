package main

import (
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/controller"
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/data"
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/model/repo"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conn := data.Connect()
	userRepo := repo.NewMysqlUserRepo(conn)
	userController := controller.NewUserController(userRepo)

	app := fiber.New()
	app.Post("user", userController.CreateUser)
	app.Listen("localhost:157")
}
