package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

var (
	config = oauth2.Config{
		ClientID:     "152904262856-jab3ep02vvgajdc66ab7ehafenoluln4.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		// This points to our Authorization Server
		// if our Client ID and Client Secret are valid
		// it will attempt to authorize our user
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
	}
)

var (
	randState = fmt.Sprintf("st%d", time.Now().UnixNano())
)
// LoadENV
func ConnectENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file not loaded properly")
	}
}
// Homepage
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage Hit!")
	u := config.AuthCodeURL(randState)
	http.Redirect(w, r, u, http.StatusFound)
}

// Authorize
func Authorize(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	state := r.Form.Get("state")
	if state != randState {
		http.Error(w, "State invalid", http.StatusBadRequest)
		return
	}

	code := r.Form.Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(*token)
}

func main() {
	ConnectENV()
	// 1 - We attempt to hit our Homepage route
	// if we attempt to hit this unauthenticated, it
	// will automatically redirect to our Auth
	// server and prompt for login credentials
	http.HandleFunc("/", HomePage)

	// 2 - This displays our state, code and
	// token and expiry time that we get back
	// from our Authorization server
	http.HandleFunc("/oauth2", Authorize)

	// 3 - We start up our Client on port 9094
	log.Println("Client is running at 9094 port.")
	log.Fatal(http.ListenAndServe(os.ExpandEnv(":${PORT}"), nil))
}
