package main

import (
	"golang.org/x/oauth2"
	"net/http"
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
