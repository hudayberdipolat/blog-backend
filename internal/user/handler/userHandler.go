package handler

import "github.com/gofiber/fiber/v2"

type handler struct{}

func NewHandler() handler {
	return handler{}
}

// user Auth begin
func (h handler) Register(ctx *fiber.Ctx) error {

	return nil
}

func (h handler) Login(ctx *fiber.Ctx) error {

	return nil
}

// user Auth end

func (h handler) GetUser(ctx *fiber.Ctx) error {

	return nil
}

func (h handler) GetAllUsers(ctx *fiber.Ctx) error {

	return nil
}

func (h handler) CreateUser(ctx *fiber.Ctx) error {

	return nil
}

func (h handler) UpdateUser(ctx *fiber.Ctx) error {

	return nil
}

func (h handler) DeleteUser(ctx *fiber.Ctx) error {

	return nil
}

func (h handler) ChangeUserPassword(ctx *fiber.Ctx) error {

	return nil
}
