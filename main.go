package main

import (
	"log"

	"github.com/rof20004/clean-go-api/pkg/account"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func main() {
	app := fiber.New()

	database := make(map[uuid.UUID]*account.Account)

	accountMemory := account.NewMemory(database)
	accountService := account.NewService(accountMemory)
	accountHandler := account.NewGofiberHandler(accountService)

	app.Get("/accounts", accountHandler.GofiberListAccounts)
	app.Post("/accounts", accountHandler.GofiberRegister)

	if err := app.Listen(":8080"); err != nil {
		log.Fatalln(err)
	}
}
