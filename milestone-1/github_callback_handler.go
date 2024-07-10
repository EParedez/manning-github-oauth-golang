package main

import (
	"context"
	"net/http"
)

func githubCallbackHandler(w http.ResponseWriter, req *http.Request) {

	// Verify if the `state` in the request parameter is one the same
	// as the `OAuthState` cookie value
	cookie := getCookie(w, req, "OAuthState")
	cookieValue := ""
	if cookie != nil {
		cookieValue = cookie.Value
	}

	if IsEmpty(cookieValue) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// If the check succeeds, request an access token from GitHub by
	// calling the oauth config's Exchange() method
	code := req.URL.Query().Get("code")
	if IsEmpty(code) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	accessToken, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		return
	}

	// Use the access token obtained to make a request to the GitHub API
	// to get the authenticated user's details

	// Create a session in the in memory session store using a random string
	// as the session identifier

	// Create a response cookie with name `Session` and value as the session identifier
	// with the maximum age set to 24 hours

	// Create a response cookie with name `OAuthState`, value as an empty string
	// with the maximum age set to -1

	// redirect the user to the URL "/"
}
