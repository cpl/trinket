package main

import "net/http"

func checkHeaders(r *http.Request) bool {
	if r.Header.Get("Content-Type") != "application/xml" {
		return false
	}

	if r.Header.Get("Accept") != "application/xml" {
		return false
	}

	return true
}

func checkAuth(username, password string) bool {
	p, ok := usersMap[username] // get user's password and if user exists
	return ok && p == password  // return if user exists and password matches
}

// remove item with index i from array s
// https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-array-in-golang
// how nice that Go does not have this built-in
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
