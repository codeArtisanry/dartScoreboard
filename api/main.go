package main

import (
	"dartscoreboard/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
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

var client *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
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
		// var user models.User
		cookie := c.Cookies("user")
		fmt.Println(cookie)
		token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("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"), nil
		})
		// claims := token.Claims.(*jwt.StandardClaims)

		// jwtware.New(jwtware.Config){
		// SigningKey: []byte("secret"),
		// })
		token = jwt.New(jwt.SigningMethodRS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["identity"] = "identity"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		fmt.Println(token)
		t, err := token.SignedString([]byte("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"))
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
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
		return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
		// Else return 403 unauthorized
	})

	// 2. Initiate google signin flow
	app.Get("/auth/:provider", func(ctx *fiber.Ctx) error {
		// TODO: Check cookie is exist [USER IS ALREADY EXIST]
		gf.BeginAuthHandler(ctx)
		// var data map[string]string
		// if err := ctx.BodyParser(&data); err != nil {
		// 	return err
		// }

		// var user models.User
		// // database.DB.Where("email = ?", data["email"]).First(&user)

		// if user.Id == 0 {
		// 	ctx.Status(fiber.StatusNotFound)
		// 	return ctx.JSON(fiber.Map{
		// 		"message": "user not found",
		// 	})
		// }

		// claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		// 	Issuer:    strconv.Itoa(int(user.Id)),
		// 	ExpiresAt: time.Now().Add(time.Hour * 42).Unix(), //2 day
		// })

		// token, err := claims.SignedString([]byte("key"))

		// if err != nil {
		// 	ctx.Status(fiber.StatusInternalServerError)
		// 	return ctx.JSON(fiber.Map{
		// 		"message": "could not login",
		// 	})
		// }

		// cookie := fiber.Cookie{
		// 	Name:     "user",
		// 	Value:    token,
		// 	Expires:  time.Now().Add(time.Hour * 24),
		// 	HTTPOnly: true,
		// }

		// ctx.Cookie(&cookie)
		// return ctx.Status(201).JSON(&fiber.Map{
		// 	"success": true,
		// 	"message": "Sign In Success",
		// })
		// IF EXIST RETURN 2xx
		// ELSE INITIATE GLAUTH SIGNIN FLOW
		return nil
	})

	log.Fatal(app.Listen(os.ExpandEnv(":${PORT}")))

}
