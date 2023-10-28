package controller

import (
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = oauth2.Config{
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

var RandomString = "random-string"

func Signin(w http.ResponseWriter, r *http.Request) {

}
