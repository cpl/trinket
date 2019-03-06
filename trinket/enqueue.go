package main

import (
	"io/ioutil"
	"net/http"
)

func handleEnqueue(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPut:
		// check headers
		// TODO define behavior for missing headers
		checkHeaders(r)

		// read body data
		bodyData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		// append to queue
		AppendToQueue(bodyData)

	case http.MethodGet:
		// TODO check auth

	}
}
