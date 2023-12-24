package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	CreateUser(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
}
