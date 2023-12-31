package main

import "net/http"

func githubCallbackHandler(w http.ResponseWriter, req *http.Request) {

	// Verify if the `state` in the request parameter is one the same
	// as the `OAuthState` cookie value

	// If the check succeeds, request an access token from GitHub by
	// calling the oauth config's Exchange() method

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
