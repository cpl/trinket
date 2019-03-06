package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
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

	switch r.Method {
	case http.MethodPut:
		// check headers
		handleHeaders(w, r)

		// ready body data
		bodyData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		// get request type
		var nodeName xml.Name
		if err := xml.Unmarshal(bodyData, &nodeName); err != nil {
			panic(err)
		}

		// validate and handle request
		switch nodeName.Local {
		case REQUEST_AVAILABILITY:
			var req RequestAvailability

			if err := xml.Unmarshal(bodyData, &req); err != nil {
				panic(err)
			}

			log.Printf("parsed request %d for %s\n",
				req.ID, nodeName.Local)

			break
		case REQUEST_BOOKINGS:
			var req RequestAvailability

			if err := xml.Unmarshal(bodyData, &req); err != nil {
				panic(err)
			}

			log.Printf("parsed request %d for %s\n",
				req.ID, nodeName.Local)

			break
		case REQUEST_CANCEL:
			var req RequestAvailability

			if err := xml.Unmarshal(bodyData, &req); err != nil {
				panic(err)
			}

			log.Printf("parsed request %d for %s\n",
				req.ID, nodeName.Local)

			break
		case REQUEST_RESERVE:
			var req RequestAvailability

			if err := xml.Unmarshal(bodyData, &req); err != nil {
				panic(err)
			}

			log.Printf("parsed request %d for %s\n",
				req.ID, nodeName.Local)

			break
		default:
			log.Printf("invalid request %s\n", nodeName.Local)
		}

	case http.MethodGet:
	}
}
