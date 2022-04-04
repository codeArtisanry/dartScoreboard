package main

import (
	"dartscoreboard/middleware"
	"dartscoreboard/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

// LoadENV
func ConnectENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file not loaded properly")
	}
}

func main() {
	ConnectENV()
	goth.UseProviders(
		google.New(os.ExpandEnv("${CLIENT_KEY}"), os.ExpandEnv("${SECRET_KEY}"), os.ExpandEnv("${PROTOCOL}://${HOST}:${PORT}/auth/google/callback"), "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"))
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Use(middleware.Validate())

	routes.Setup(app)

	log.Fatal(app.Listen(os.ExpandEnv(":${PORT}")))
}
