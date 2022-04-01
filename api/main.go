package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	gf "github.com/shareed2k/goth_fiber"
)

type Cookie struct {
	Name     string
	Value    string
	Path     string
	Domain   string
	Expires  time.Time
	Secure   bool
	HTTPOnly bool
	SameSite string // lax, strict, none
}

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
		google.New(os.ExpandEnv("${CLIENT_KEY}"), os.ExpandEnv("${SECRET_KEY}"), os.ExpandEnv("${PROTOCOL}://${HOST}:${PORT}/auth/google/callback")))
	app := fiber.New()

	// 1. Endpoint for i am logged in?
	app.Get("/", func(c *fiber.Ctx) error {
		// TODO: Check from cookie if user is exist or not
		c.Redirect("/auth/google")
		// If exist then return 2xx status code
		// Else return 403 unauthorized
		return c.JSON(fiber.Map{"status": "Fail", "message": "unauthorized user"})
	})

	// 2. Initiate google signin flow
	app.Get("/auth/:provider", func(ctx *fiber.Ctx) error {
		// TODO: Check cookie is exist [USER IS ALREADY EXIST]
		gf.BeginAuthHandler(ctx)

		// IF EXIST RETURN 2xx
		// ELSE INITIATE GLAUTH SIGNIN FLOW
		return nil
	})

	// 3. Redirect by google
	app.Get("/auth/:provider/callback", func(ctx *fiber.Ctx) error {
		userinfo, err := gf.CompleteUserAuth(ctx)
		if err != nil {
			return err
		}
		// GET TOKEN
		fmt.Println("ID TOKEN", userinfo.IDToken)

		// TODO: Set cookie
		cookie := new(fiber.Cookie)
		cookie.Name = "userinfo"
		cookie.Value = userinfo.IDToken
		cookie.Expires = time.Now().Add(30 * time.Hour * 24)
		cookie.HTTPOnly = true
		// Set cookie from JWT
		ctx.Cookie(cookie)
		// TODO: Redirect user to frontend
		ctx.Redirect("/home")
		return ctx.JSON(fiber.Map{
			"message": "success",
			"data":    userinfo,
		})
	})

	// 4. Signout
	app.Get("/logout/:provider", func(ctx *fiber.Ctx) error {
		gf.Logout(ctx)
		// Clear all cookie
		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		}

		ctx.Cookie(&cookie)

		return ctx.JSON(fiber.Map{
			"message": "success",
		})
		// Return 200
	})

	//frontend ROUTE
	app.Get("/home", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "successfully in homepage",
		})
	})
	log.Fatal(app.Listen(os.ExpandEnv(":${PORT}")))

}
