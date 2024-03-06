package main

import (
	"fmt"
	"os"

	"github.com/Podre-Henrique/arquitetura-api/mvc/api/controller"
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/data"
	"github.com/Podre-Henrique/arquitetura-api/mvc/api/model/repo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var (
	PORT = os.Getenv("PORT")
)

func main() {
	fmt.Println("Servidor iniciado")

	conn := data.Connect()
	userRepo := repo.NewMysqlUserRepo(conn)
	userController := controller.NewUserController(userRepo)

	app := fiber.New()
	app.Post("user", userController.CreateUser)

	app.Listen(":" + PORT)
}
