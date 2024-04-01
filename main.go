package main

import (
	"fmt"
	"os"

	"github.com/Anishgupta26/user-app/health"
	"github.com/Anishgupta26/user-app/users"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

// root check route
func root(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"Hello": "World"})
}

func main() {
	//creating fiber app
	app := fiber.New()
	app.Get("/", root)
	app.Get("/health", health.Check)
	app.Get("/userinfo", users.User_func)
	app.Post("/post", func(c *fiber.Ctx) error {
		body := c.Body()
		res := string(body)
		fmt.Println(res)
		return c.SendString(res)
	})
	//Assignment-
	app.Post("/user", users.Add_user)

	err := app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		logrus.Fatalln(err)
		return
	}
	fmt.Print(err)
	// POST /user  CREATE
	/*
		Body
		{
			"first_name":"anish",
			"last_name":"gupta"
		}

		Response:
		{
			"response": "user created",
			"error":""
		}
		STatus: 200
		Body
		{
			"first_name":200,
			"last_name":"gupta"
		}

		Response:
		{
			"response": "",
			"error":"first_name can't be integer"
		}
		Status: 400
	*/

	// GET /user FETCH ALL USERS

	// GET /user?first_name=anish FETCH SINGLE USER

}
