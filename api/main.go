package main

import (
	models "dartscoreboard/models/database"
	"dartscoreboard/routes"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	models.Migration("dart.db")
	ConnectENV()
	goth.UseProviders(
		google.New(os.ExpandEnv("${CLIENT_KEY}"), os.ExpandEnv("${SECRET_KEY}"), os.ExpandEnv("${PROTOCOL}://${HOST}:${AUTHPORT}/auth/google/callback"), "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"))
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	app.Use(logger.New())
	routes.Setup(app)
	// app.Group("/", middleware.Validate())
	// app.Use(middleware.Validate())

	log.Fatal(app.Listen(os.ExpandEnv(":${PORT}")))
}
