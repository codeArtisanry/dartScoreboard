package main

import (
	"dartscoreboard/models"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	gf "github.com/shareed2k/goth_fiber"
)

const (
	key = "152904262856-jab3ep02vvgajdc66ab7ehafenoluln4.apps.googleusercontent.com"
	sec = "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"
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
	models.Database()
	goth.UseProviders(
		google.New(key, sec, "http://localhost:3000/auth/google/callback"),
	)

	app := fiber.New()

	app.Get("/auth/:provider/callback", func(ctx *fiber.Ctx) error {
		user, err := gf.CompleteUserAuth(ctx)
		if err != nil {
			return err
		}
		ctx.JSON(user)
		return nil
	})

	app.Get("/logout/:provider", func(ctx *fiber.Ctx) error {
		gf.Logout(ctx)
		ctx.Redirect("/")
		return nil
	})

	app.Get("/auth/:provider", func(ctx *fiber.Ctx) error {
		if gothUser, err := gf.CompleteUserAuth(ctx); err == nil {
			ctx.JSON(gothUser)
		} else {
			gf.BeginAuthHandler(ctx)
		}
		return nil
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Format("<p><a href='/auth/google'>google</a></p>")
		return nil
	})

	log.Fatal(app.Listen(os.ExpandEnv(":${PORT}")))

}
