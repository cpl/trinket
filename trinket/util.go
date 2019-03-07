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
