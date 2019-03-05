package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleHeaders(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/xml" {
		w.Write([]byte("Missing Content-Type header\n"))
	}

	if r.Header.Get("Accept") != "application/xml" {
		w.Write([]byte("Missing Accept header\n"))
	}
}

func handleEnqueue(w http.ResponseWriter, r *http.Request) {

	handleHeaders(w, r)

	switch r.Method {
	case http.MethodPut:
		// ready body data
		bodyData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		var req RequestReserve
		if err := xml.Unmarshal(bodyData, &req); err != nil {
			panic(err)
		}

		fmt.Println(req)

	case http.MethodGet:
	}
}

func handleBookings(w http.ResponseWriter, r *http.Request) {

}

func handleAvailable(w http.ResponseWriter, r *http.Request) {

}
