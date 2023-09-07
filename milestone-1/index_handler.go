package main

import "net/http"

func indexHandler(w http.ResponseWriter, req *http.Request) {
	// Check for session HTTTP cookie in the request

	// If cookie present and it's value a valid session identifier,
	// return successful response

	// if cookie not found or the value cannot be found in the sessionStore
	// initiate GitHub authentication process
}
