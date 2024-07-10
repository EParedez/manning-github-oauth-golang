package main

import "strings"

func IsEmpty(str string) bool {
	return strings.Trim(str, " ") == ""
}
