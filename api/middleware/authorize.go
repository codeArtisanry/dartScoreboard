package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jellydator/ttlcache/v2"
)

var cache ttlcache.SimpleCache = ttlcache.NewCache()

// Get Public Key into the Cache Memory and Validate Token With that Public Key
func Validate(Vaconfig ...fiber.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var publicKey string = "publicKey"
		cookie := ctx.Cookies("user")
		if cookie == "" {
			return ctx.Redirect("/auth/google")
		}
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
			// Get Public key from cache memory
			pem, err := cache.Get(publicKey)
			if err != nil {
				pem, err = getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
				if err != nil {
					return nil, err
				}
			}
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem.(string)))
			if err != nil {
				// If get error to fetch a public key than set again public key in cache memory
				pem, err = getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
				if err != nil {
					return nil, err
				}
			}
			return key, nil
		})
		ctx.Locals("claims", claims["email"])
		fmt.Println(token)
		// ... error handling
		if err != nil {
			fmt.Println("err", err)
			return ctx.Redirect("/auth/google")
		} else {
			ctx.Redirect(os.ExpandEnv("${PROTOCOL}://${HOST}:${FRONTENDPORT}/"))
			return ctx.Next()
		}
	}
}

// Google Public key
func getGooglePublicKey(keyID string) (string, error) {
	var publicKey string = "publicKey"
	cache.SetTTL(time.Duration(24 * time.Hour))
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
	publicKeyValue := key
	// Set public key in cache memory
	cache.Set(publicKey, publicKeyValue)
	return key, nil
}
