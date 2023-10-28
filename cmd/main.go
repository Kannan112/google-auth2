package main

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = oauth2.Config{
		ClientID:     "744537701078-glv87e3sdrgfrqg74na2a15b6onj8tca.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/admob.readonly",
			"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/cloud-platform.read-only",
		},
		Endpoint: google.Endpoint,
	}
)

func main() {

	http.HandleFunc("/", InitialPageHandler)
	http.HandleFunc("/auth/google/login", GoogleLoginHandler)
	http.HandleFunc("/auth/google/callback", GoogleCallbackHandler)

	http.ListenAndServe(":8080", nil)
}

func InitialPageHandler(w http.ResponseWriter, r *http.Request) {

}

var RandomString = "random-string"

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	// Exchange the code for an access token
	token, err := googleOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		// Handle the error
	}
	fmt.Println(token)

	// Use the token to make requests to the Google API or fetch user info

	// Redirect or show the user's profile page
}
