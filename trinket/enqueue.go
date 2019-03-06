package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

func handleEnqueue(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPut:
		// check headers
		checkHeaders(r)

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
		// --------------------------------------------------------------------
		case REQUEST_AVAILABILITY:
			var req RequestAvailability

			if err := xml.Unmarshal(bodyData, &req); err != nil {
				panic(err)
			}

			if !checkAuth(req.Username, req.Password) {
				log.Printf("invalid auth for request %d\n", req.ID)
				return
			}

			var res ResponseAvailability

			res.XMLName = xml.Name{Local: "response"}
			res.Code = 200
			for idx := range slots {
				if slots[idx] == "free" {
					res.Slots = append(res.Slots, idx)
				}
			}

			xmlResponse, err := xml.Marshal(res)
			if err != nil {
				panic(err)
			}

			if _, err := w.Write(xmlResponse); err != nil {
				panic(err)
			}

			log.Printf("parsed request %d for %s\n",
				req.ID, nodeName.Local)

			break
		// --------------------------------------------------------------------
		case REQUEST_BOOKINGS:
			var req RequestAvailability

			if err := xml.Unmarshal(bodyData, &req); err != nil {
				panic(err)
			}

			if !checkAuth(req.Username, req.Password) {
				log.Printf("invalid auth for request %d\n", req.ID)
				return
			}

			log.Printf("parsed request %d for %s\n",
				req.ID, nodeName.Local)

			break
		// --------------------------------------------------------------------
		case REQUEST_CANCEL:
			var req RequestAvailability

			if err := xml.Unmarshal(bodyData, &req); err != nil {
				panic(err)
			}

			if !checkAuth(req.Username, req.Password) {
				log.Printf("invalid auth for request %d\n", req.ID)
				return
			}

			log.Printf("parsed request %d for %s\n",
				req.ID, nodeName.Local)

			break
		// --------------------------------------------------------------------
		case REQUEST_RESERVE:
			var req RequestAvailability

			if err := xml.Unmarshal(bodyData, &req); err != nil {
				panic(err)
			}

			if !checkAuth(req.Username, req.Password) {
				log.Printf("invalid auth for request %d\n", req.ID)
				return
			}

			log.Printf("parsed request %d for %s\n",
				req.ID, nodeName.Local)

			break
		// --------------------------------------------------------------------
		default:
			log.Printf("invalid request %s\n", nodeName.Local)
		}

	case http.MethodGet:
	}
}
