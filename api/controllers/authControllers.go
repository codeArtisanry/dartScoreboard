package controllers

import (
	models "dartscoreboard/models/database"
	types "dartscoreboard/models/types"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	gf "github.com/shareed2k/goth_fiber"
)

var (
	db = models.Database("dart.db")
)

func Endpoint(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("user")
	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	if err != nil {
		fmt.Println("err  :", err)
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "User NOT Exist",
		})
		return ctx.Redirect("/auth/google")
	} else {
		ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"message": "User Exist",
		})
		return ctx.Redirect(os.ExpandEnv("${PROTOCOL}://${HOST}:${FRONTENDPORT}/"))
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
	UserInfo := types.User{
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
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.HTTPOnly = false
	cookie.SameSite = fiber.CookieSameSiteNoneMode
	// Set cookie from JWT
	ctx.Cookie(cookie)
	// TODO: Redirect user to frontend
	return ctx.Redirect(os.ExpandEnv("${PROTOCOL}://${HOST}:${FRONTENDPORT}/"))
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
	return ctx.Redirect(os.ExpandEnv("${PROTOCOL}://${HOST}:${FRONTENDPORT}/${SIGNOUT}"))
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
