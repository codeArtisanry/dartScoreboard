package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Validate(config ...fiber.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		cookie := ctx.Cookies("user")
		if cookie == "" {
			return ctx.Redirect("/auth/google")
		}
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
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
		ctx.Locals("claims", claims["email"])
		fmt.Println(token)
		// ... error handling
		if err != nil {
			fmt.Println("err",err)
			return ctx.Redirect("auth/google")
		} else {
			ctx.Redirect(os.ExpandEnv("${PROTOCOL}://${HOST}:${FRONTENDPORT}/"))
			return ctx.Next()
		}
	}
}

// Google Public key
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
