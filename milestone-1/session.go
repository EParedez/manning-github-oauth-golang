package main

type userData struct {
	Login       string
	AccessToken string
}

var sessionStore map[string]userData
