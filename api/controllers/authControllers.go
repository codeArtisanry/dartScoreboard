package controllers

import (
	"dartscoreboard/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	gf "github.com/shareed2k/goth_fiber"
)

var db = models.Database()

// 1. Endpoint for i am logged in?
func Endpoint(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("user")
	fmt.Println("this is cookie", cookie)
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
		if err != nil {
			return nil, err
		}
		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, err
		}
		return key, nil
	})
	claims := token.Claims.(*jwt.StandardClaims)
	fmt.Println("claims :", claims)
	if err != nil {
		fmt.Println("err  :", err)
		// return ctx.Redirect("/auth/google")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User NOT Exist",
		})
	} else {
		// return ctx.Redirect("/home")
		return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message": "User Exist",
		})
	}
	// If exist then return 2xx status code
	// Else return 403 unauthorized
}

// 2. Initiate google signin flow
func Signinflow(ctx *fiber.Ctx) error {
	// TODO: Check cookie is exist [USER IS ALREADY EXIST]
	gf.BeginAuthHandler(ctx)

	// IF EXIST RETURN 2xx
	// ELSE INITIATE GLAUTH SIGNIN FLOW
	return nil
}

// 3. Redirect by google
func GoogleRedirect(ctx *fiber.Ctx) error {
	user, err := gf.CompleteUserAuth(ctx)
	if err != nil {
		return err
	}
	// fmt.Println(user)
	UserInfo := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
	}
	id, err := models.VerifyAndInsertUser(db, UserInfo)
	fmt.Println(id)
	if err != nil {
		fmt.Println(err)
	} // GET TOKEN &
	// TODO: Set cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "user"
	cookie.Value = user.IDToken
	cookie.Expires = time.Now().Add(30 * time.Hour * 24)
	cookie.HTTPOnly = false
	cookie.SameSite = fiber.CookieSameSiteNoneMode
	// Set cookie from JWT
	ctx.Cookie(cookie)
	// TODO: Redirect user to frontend
	endPointURL := fmt.Sprintf(os.ExpandEnv("${PROTOCOL}://${HOST}:${FRONTENDPORT}/${REDIRECTPOINT}"))
	return ctx.Redirect(endPointURL)
	// return ctx.Redirect(os.ExpandEnv("${PROTOCOL}://${HOST}:${FRONTENDPORT}/home"))
	// ctx.Redirect("/home/:id")

}

// 4. Signout
func Signout(ctx *fiber.Ctx) error {
	gf.Logout(ctx)
	// Clear all cookie
	cookie := fiber.Cookie{
		Name:     "user",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: false,
	}

	ctx.Cookie(&cookie)

	// Return 200
	return ctx.Redirect(os.ExpandEnv("${PROTOCOL}://${HOST}:${FRONTENDPORT}"))
	// return ctx.JSON(fiber.Map{
	// 	"message": "success",
	// })
}

func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}
	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}

func UserJson(ctx *fiber.Ctx) (goth.User, error) {
	user, err := gf.CompleteUserAuth(ctx)
	if err != nil {
		return user, err
	}
	return user, err
}
