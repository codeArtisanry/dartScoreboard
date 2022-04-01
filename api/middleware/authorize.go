package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

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
	fmt.Println("this is claims", claims)
	if err != nil {
		fmt.Println("err  :", err)
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "User NOT Exist",
		})
	} else {
		ctx.Status(fiber.StatusAccepted)
		return ctx.JSON(fiber.Map{
			"message": "User Exist",
		})
	}
	// c.Redirect("/auth/google")
	// If exist then return 2xx status code
	// Else return 403 unauthorized
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
