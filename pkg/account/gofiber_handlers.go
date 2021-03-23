package account

import (
	"github.com/gofiber/fiber/v2"
)

type GofiberHandler struct {
	serv *Service
}

func NewGofiberHandler(serv *Service) GofiberHandler {
	return GofiberHandler{serv}
}

func (gf GofiberHandler) GofiberRegister(c *fiber.Ctx) error {
	var account Account
	if err := c.BodyParser(&account); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := gf.serv.Register(&account); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(account)
}

func (gf GofiberHandler) GofiberListAccounts(c *fiber.Ctx) error {
	accounts, err := gf.serv.ListAccounts()
	if err != nil && err != ErrNoAccountsFound {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if err != nil && err == ErrNoAccountsFound {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.JSON(accounts)
}
