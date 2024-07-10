package main

import (
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"log"
	"net/http"
	"os"
)

var oauthConf *oauth2.Config

func main() {

	oauthConf.Scopes = []string{"repo", "user"}
	oauthConf.Endpoint = oauth2.Endpoint{
		AuthURL: "",
	}

	server := http.Server{
		Addr: "http://localhost:8080",
	}
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/github/oauth/", githubCallbackHandler)
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func getCookie(w http.ResponseWriter, r *http.Request, cookieName string) *http.Cookie {
	c, err := r.Cookie(cookieName)
	if err != nil {
		fmt.Fprintln(w, "Cannot get the cookie")
	}
	return c
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: false,
	}

	http.SetCookie(w, &c1)
}

func initOAuthConfig() {
	if len(os.Getenv("CLIENT_ID")) == 0 || len(os.Getenv("CLIENT_SECRET")) == 0 {
		log.Fatal("Must specify your app's CLIENT_ID and CLIENT_SECRET")
	}

	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"repo", "user"},
		Endpoint:     github.Endpoint,
	}
}
