package controllers

import (
	"dartscoreboard/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
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
	userinfo := models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
	}
	// db := models.Database()
	fmt.Println("from api", userinfo)
	// models.InsertData(db, userinfo)

	// GET TOKEN &
	// TODO: Set cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "user"
	cookie.Value = user.IDToken
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
}

// 4. Signout
func Signout(ctx *fiber.Ctx) error {
	gf.Logout(ctx)
	// Clear all cookie
	cookie := fiber.Cookie{
		Name:     "user",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	// Return 200
	return ctx.JSON(fiber.Map{
		"message": "success",
	})
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
