package user

import (
	"log"

	userService "myFirstApiWithGo/user/service"

	structUser "myFirstApiWithGo/user/struct"

	"github.com/gofiber/fiber/v2"
)

func UserController(app *fiber.App) {
	route := app.Group("/user")
	//se crea la app de fiber
	route.Get("/get-user", func(c *fiber.Ctx) error {
		user := c.Query("user")
		log.Println("entro al get", user)

		return c.JSON(userService.GetUser(user))
	})
	route.Post("/create-user", func(c *fiber.Ctx) error {
		log.Println("entro al post")
		user := new(structUser.User)
		log.Println("entro al post")
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		log.Println("entro al post", user)

		return c.JSON(userService.CreateUser(user))
	})
	route.Put("/upadate-user", func(c *fiber.Ctx) error {
		user := new(structUser.User)

		if err := c.BodyParser(&user); err != nil {
			return err
		}
		// log.Println(user)
		return c.JSON(userService.UpdateUser(user))
	})
	route.Delete("/delete-user", func(c *fiber.Ctx) error {
		user := new(structUser.User)
		if err := c.BodyParser(&user); err != nil {
			return err
		}
		// log.Println(user)
		return c.JSON(userService.DeleteUser(user.Name))

	})
	//esto hace que corra la route

}
