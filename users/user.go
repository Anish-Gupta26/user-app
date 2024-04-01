package users

import (
	_ "encoding/json"
	"fmt"
	_ "reflect"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func User_func(c *fiber.Ctx) error {

	data := map[string]string{
		"first_name": "ANish",
		"last_name":  "Gupta",
	}
	fmt.Print(c.Query("name"))
	return c.JSON(data)
}

func Add_user(c *fiber.Ctx) error {
	var user User
	err := c.BodyParser(&user)
	fmt.Println(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	fmt.Println(err)
	return c.JSON(user)
}

// func Add_user(c *fiber.Ctx) error {
// 	payload := string(c.Body())
// 	var data map[string]interface{}
// 	err := json.Unmarshal([]byte(payload), &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	nameResponseType := reflect.TypeOf(data["name"]).Kind()
// 	if nameResponseType == reflect.String {
// 		responseData := map[string]interface{}{"response": "user created", "error": ""}

// 		return c.JSON(responseData)
// 	} else {
// 	}
// 	return c.SendString("post unsuccessful")
// }
