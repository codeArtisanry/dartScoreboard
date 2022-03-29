package main

import (
	"dartscoreboard/models"
	"fmt"
	"log"
	"os"
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
		google.New(os.ExpandEnv("${CLIENT_KEY}"), os.ExpandEnv("${SECRET_KEY}"), os.ExpandEnv("${PROTOCOL}://${HOST}:${PORT}/auth/google/callback")))
	app := fiber.New()

	// 1. Endpoint for i am logged in?
	app.Get("/", func(c *fiber.Ctx) error {
		// TODO: Check from cookie if user is exist or not
		// if exist then return 2xx status code
		if c.Cookies("userinfo") == c.Cookies("userinfo") {
			fmt.Println(c.Cookies("userinfo"))
			return c.JSON(fiber.Map{"status": "success", "message": "Success login"})
			// Else return 403 unauthorized
		} else {
			c.Redirect("/auth/google")
		}
		// If exist then return 2xx status code
		// return nil
		return c.JSON(fiber.Map{"status": "Fail", "message": "unauthorized user"})
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
		////////////////////////changes by jeel/////////////////////////////////////
		user := models.User{
			Id:      userinfo.RawData["id"],
			Email:   userinfo.RawData["email"],
			Picture: userinfo.RawData["picture"],
		}
		// db := models.Database()
		// models.InsertData(db, user)
		fmt.Println("from api", user)
		////////////////////////////////////////////////////////////////////////////
		if user.Id == 0 {
			ctx.Status(fiber.StatusNotFound)
			return ctx.JSON(fiber.Map{
				"message": "user not found",
			})
		}
		fmt.Println(userinfo)
		// GET TOKEN
		// token := ctx.Cookies("userinfo")
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["identity"] = userinfo
		claims["admin"] = true
		claims["exp"] = time.Now().Add(24 * time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(`eyJhbGciOiJSUzI1NiIsImtpZCI6IjU4YjQyOTY2MmRiMDc4NmYyZWZlZmUxM2MxZWIxMmEyOGRjNDQyZDAiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIxNTI5MDQyNjI4NTYtamFiM2VwMDJ2dmdhamRjNjZhYjdlaGFmZW5vbHVsbjQuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIxNTI5MDQyNjI4NTYtamFiM2VwMDJ2dmdhamRjNjZhYjdlaGFmZW5vbHVsbjQuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMTM1MzEyMjk1MzA4MTgyNjU4ODMiLCJoZCI6ImltcHJvd2lzZWQuY29tIiwiZW1haWwiOiJ2YXRzYWxAaW1wcm93aXNlZC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6IlNvUGV4aGdoRlJ2bXNQN3FPNWZ5V0EiLCJpYXQiOjE2NDg1MzAwMTUsImV4cCI6MTY0ODUzMzYxNX0.MzHc_mGAF3tdeeklVWVlbV_8KiInfXcv_cVvz9P-2CJIEHW2HA63eO1W3VxT22sXGZkegvhFYdpMyYo_L8fugSX-4TbaRUjSOPzmQXAgoiXMfedHNGejQH50ciGxp4KZVo3P8sC4vZJglzyEh8pFOlISFjAQiIh3vWkLDZTcmAuEmzo5KlM4O88e5k7F47dS_qArqpPNRpvr1gOddf5HvYYYkRVLem2njLxc9qe7oqX80GGZW32zkUoPe4467TRsh0T9uihPj6ue8SBAN60vnEqtnrZ7c_YPjNGXJaPG3HfilxraBtWNp1Jl_8t4vSLtq9wbKHISLibXGaXLs8ykFQ`))
		fmt.Println("token by google: ", t)
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		// if err != nil {
		// 	ctx.Status(fiber.StatusInternalServerError)
		// 	return ctx.JSON(fiber.Map{
		// 		"message": "could not login",
		// 	})
		// }

		// TODO: Set cookie
		cookie := new(fiber.Cookie)
		cookie.Name = "userinfo"
		cookie.Value = t
		cookie.Expires = time.Now().Add(30 * time.Hour * 24)
		// Set cookie from JWT
		ctx.Cookie(cookie)
		// TODO: Redirect user to frontend
		ctx.Redirect("/home")
		return ctx.JSON(fiber.Map{
			"message": "success",
			"data":    t,
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

	// ///////////////////////chnages by jeel///////////////////////////////////
	//frontend ROUTE
	app.Get("/home", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "successfully in homepage",
		})
	})
	log.Fatal(app.Listen(os.ExpandEnv(":${PORT}")))

}
