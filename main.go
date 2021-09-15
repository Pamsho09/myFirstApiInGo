package main

import (
	"log"

	useController "myFirstApiWithGo/user"
 	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type User struct {
	Numero1 int `json:"numero1"`
	Numero2 int `json:"numero2"`
}

func main() {
	godotenv.Load()
	//se crea la app de fiber
	app := fiber.New()
	app.Use(cors.New())
	//crea una ruta para el home
	log.Println("Iniciando el servidor en el puerto 3000")
	useController.UserController(app)
	//esto hace que corra la app
	log.Fatal(app.Listen(":8080"))
}
