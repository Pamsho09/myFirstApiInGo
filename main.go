package main

import (
	"log"
	"os"

	useController "myFirstApiWithGo/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

type User struct {
	Numero1 int `json:"numero1"`
	Numero2 int `json:"numero2"`
}
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }
func main() {
	godotenv.Load()
	//se crea la app de fiber
	app := fiber.New()
	app.Use(cors.New())
	//crea una ruta para el home
	log.Println("Iniciando el servidor en el puerto ",goDotEnvVariable("PORT"))
	useController.UserController(app)
	//esto hace que corra la app
	log.Fatal(app.Listen(":"+goDotEnvVariable("PORT")))
}
