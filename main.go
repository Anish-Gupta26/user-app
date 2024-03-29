package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type Person struct {
	Name string `json:"name"`
}

func health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"Hello": "World"})
}

func main() {
	app := fiber.New()
	app.Get("/", health)
	err := app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		logrus.Fatalln(err)
		return
	}

}
