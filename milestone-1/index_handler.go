package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	// Check for session HTTTP cookie in the request
	cookie := getCookie(w, req, "Session")
	cookieValue := ""
	if cookie != nil {
		cookieValue = cookie.Value
	}

	// If cookie present and it's value a valid session identifier,
	// return successful response
	if val, ok := sessionStore[cookieValue]; ok {
		name := val.Login
		w.Write([]byte(fmt.Sprintf("Successfully authorized to access GitHub on your behalf: %s", name)))
	}
	// if cookie not found or the value cannot be found in the sessionStore
	// initiate GitHub authentication process
	url := oauthConf.AuthCodeURL("statedummy")
	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
	return
}
