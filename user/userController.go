package user

import (
	"log"

	userService "myFirstApiWithGo/user/service"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func UserController(app *fiber.App) {
	route := app.Group("/user")
	//se crea la app de fiber
	route.Get("/get-user", func(c *fiber.Ctx) error {
		log.Println("entro al get")
		return c.JSON(userService.GetUser("user"))
	})
	route.Post("/create-user", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(&user); err != nil {
			return err
		}
		log.Println(user)
		return c.JSON("hola")
	})
	route.Put("/upadate-user", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(&user); err != nil {
			return err
		}
		// log.Println(user)
		return c.JSON("hp;a")
	})
	route.Delete("/delete-user", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		// log.Println(user)
		return c.JSON("hp;a")

	})
	//esto hace que corra la route

}
