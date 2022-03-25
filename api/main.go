package main

import (
	"dartscoreboard/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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
		google.New(os.ExpandEnv("${CLIENT_KEY}"), os.ExpandEnv("${SECRET_KEY}"), os.ExpandEnv("${PROTOCOL}://${HOST}:${PORT}/auth/google/callback")),
	)

	app := fiber.New()

	// 3. Redirect by google
	app.Get("/auth/:provider/callback", func(ctx *fiber.Ctx) error {
		user, err := gf.CompleteUserAuth(ctx)
		if err != nil {
			return err
		}

		// TODO: Set cookie
		cookie := new(fiber.Cookie)
		cookie.Name = "user"
		cookie.Value = user.IDToken
		cookie.Expires = time.Now().Add(30 * time.Hour * 24)
		// Set cookie
		ctx.Cookie(cookie)
		// TODO: Redirect user to frontend
		ctx.Redirect("/")
		return err
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

	// 1. Endpoint for i am logged in?
	app.Get("/", func(c *fiber.Ctx) error {
		// TODO: Check from cookie if user is exist

		cookie := c.Cookies("jwt")
		_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.ExpandEnv("${SECRET_KEY}")), nil
		})
		// fmt.Println(token)

		// claims := token.Claims.(*jwt.StandardClaims)
		var user models.User
		// userData := models.User{Token: claims}
		// fmt.Println(userData)
		// fmt.Println(claims)
		// database.DB.Where("id = ?", claims.Issuer).First(&user)
		if err != nil {
			fmt.Println(err)
			c.Status(fiber.StatusUnauthorized)
			return c.JSON(fiber.Map{
				"message": "unauthenticated",
			})
		}
		// c.Redirect("/auth/google")
		// If exist then return 2xx status code
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "Login successfully",
			"users":   user,
		})
		// Else return 403 unauthorized
	})

	// 2. Initiate google signin flow
	app.Get("/auth/:provider", func(ctx *fiber.Ctx) error {
		// TODO: Check cookie is exist [USER IS ALREADY EXIST]
		gf.BeginAuthHandler(ctx)
		var data map[string]string
		if err := ctx.BodyParser(&data); err != nil {
			return err
		}

		var user models.User
		// database.DB.Where("email = ?", data["email"]).First(&user)

		if user.Id == 0 {
			ctx.Status(fiber.StatusNotFound)
			return ctx.JSON(fiber.Map{
				"message": "user not found",
			})
		}

		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(int(user.Id)),
			ExpiresAt: time.Now().Add(time.Hour * 42).Unix(), //2 day
		})

		token, err := claims.SignedString([]byte(os.ExpandEnv("${SECRET_KEY}")))

		if err != nil {
			ctx.Status(fiber.StatusInternalServerError)
			return ctx.JSON(fiber.Map{
				"message": "could not login",
			})
		}

		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		}

		ctx.Cookie(&cookie)
		return ctx.Status(201).JSON(&fiber.Map{
			"success": true,
			"message": "",
		})
		// IF EXIST RETURN 2xx
		// ELSE INITIATE GLAUTH SIGNIN FLOW
	})

	log.Fatal(app.Listen(os.ExpandEnv(":${PORT}")))

}
