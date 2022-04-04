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

func Validate(config ...fiber.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
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
		ctx.Redirect("/auth/google")
		if err != nil {
			fmt.Println("err  :", err)
			ctx.Redirect("/auth/google")
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "User NOT Exist",
			})
		} else {
			ctx.Redirect("/home")
			return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
				"message": "User Exist",
			})
		}
	}
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
