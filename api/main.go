package main

import (
	"dartscoreboard/models"
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
	models.Database()
	goth.UseProviders(
		google.New(os.ExpandEnv("${CLIENT_KEY}"), os.ExpandEnv("${SECRET_KEY}"), os.ExpandEnv("${PROTOCOL}://${HOST}:${PORT}/auth/google/callback"), os.ExpandEnv("[]string{${SCOPES}}")))
	fmt.Println(os.Getenv("${SCOPES}"))
	app := fiber.New()

	// 1. Endpoint for i am logged in?
	app.Get("/", func(c *fiber.Ctx) error {
		// TODO: Check from cookie if user is exist
		// var user models.User
		// userData := models.User{Token: claims}
		// fmt.Println(userData)
		// database.DB.Where("id = ?", claims.Issuer).First(&user)
		// if err != nil {
		// 	fmt.Println(err, "key error")
		// 	c.Status(fiber.StatusUnauthorized)
		// 	return c.JSON(fiber.Map{
		// 		"message": "unauthenticated",
		// 	})
		// }
		c.Redirect("/auth/google")
		// If exist then return 2xx status code
		return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": "tokens"})
		// Else return 403 unauthorized
	})

	// 2. Initiate google signin flow
	app.Get("/auth/:provider", func(ctx *fiber.Ctx) error {
		// TODO: Check cookie is exist [USER IS ALREADY EXIST]
		gf.BeginAuthHandler(ctx)

		// var user models.User

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
		var user models.User
		////////////////////////changes by jeel/////////////////////////////////////
		userData := models.User{
			Id:    userinfo.RawData["id"],
			Email: userinfo.RawData["email"],
		}
		fmt.Println(userData)
		////////////////////////////////////////////////////////////////////////////
		if user.Id == 0 {
			ctx.Status(fiber.StatusNotFound)
			return ctx.JSON(fiber.Map{
				"message": "user not found",
			})
		}
		// GET TOKEN
		token := ctx.Cookies("user")

		if err != nil {
			ctx.Status(fiber.StatusInternalServerError)
			return ctx.JSON(fiber.Map{
				"message": "could not login",
			})
		}

		fmt.Println("token by google: ", token)

		// TODO: Set cookie
		cookie := new(fiber.Cookie)
		cookie.Name = "userinfo"
		cookie.Value = token
		cookie.Expires = time.Now().Add(30 * time.Hour * 24)
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
	///////////////////////chnages by jeel///////////////////////////////////
	app.Get("/home/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "successfully in homepage",
		})
	})
	//////////////////////////////////////////////////////////////////////////

	log.Fatal(app.Listen(os.ExpandEnv(":${PORT}")))

}
