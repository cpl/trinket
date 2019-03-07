package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
)

func handleEnqueuePUT(w http.ResponseWriter, r *http.Request) {

	// chance to fail request "service is busy"
	fmt.Println(rand.Uint32() % 100)
	if rand.Uint32()%100 < failChance {
		log.Println("request failed, service unavailable")

		// wrote headers
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Header().Add("Status", "503")
		w.Header().Add("Server", "Trinket")
		w.Header().Add("Connection", "close")
		w.Header().Add("Cache-Control", "private, max-age=0, must-revalidate")

		// write 503
		w.WriteHeader(http.StatusServiceUnavailable)

		return
	}

	// check headers
	if !checkHeaders(r) {
		w.Write([]byte("missing expected headers as requested by labscript\n"))
		return
	}

	// read body data
	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	mid := AppendToQueue(bodyData)
	// append request to queue and format URI with msg id
	uri := fmt.Sprintf(
		"<msg_uri>http://localhost:%d/queue/msg/%d</msg_uri>",
		port, mid)

	// handle headers
	w.Header().Add("Content-Type", "application/xml; charset=utf-8")
	w.Header().Add("Status", "200")
	w.Header().Add("Server", "Trinket")
	w.Header().Add("Connection", "close")
	w.Header().Add("Cache-Control", "private, max-age=0, must-revalidate")

	// write response URI
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(uri))

	log.Printf("PUT request %d in queue\n", mid)
}

func handleEnqueueGET(w http.ResponseWriter, r *http.Request) {

	// chance to fail request "service is busy"
	fmt.Println(rand.Uint32() % 100)
	if rand.Uint32()%100 < failChance {
		log.Println("request failed, service unavailable")

		// wrote headers
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Header().Add("Status", "503")
		w.Header().Add("Server", "Trinket")
		w.Header().Add("Connection", "close")
		w.Header().Add("Cache-Control", "private, max-age=0, must-revalidate")

		// write 503
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service unavailable"))

		return
	}

	// parse msgID and handle invalid ID
	msgID, err := strconv.Atoi(mux.Vars(r)["msg_id"])
	if err != nil || msgID < 1 || msgID > len(listingQueue) {
		log.Printf("request failed, invalid msgID: %s\n",
			mux.Vars(r)["msg_id"])

		// wrote headers
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Header().Add("Status", "503")
		w.Header().Add("Server", "Trinket")
		w.Header().Add("Connection", "close")
		w.Header().Add("Cache-Control", "private, max-age=0, must-revalidate")

		// write 503
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service unavailable, wrong msg_id"))

		return
	}

	// handle listing not processed
	if listingQueue[msgID-1].Status != "processed" {
		log.Printf("request failed, message %d not processed\n", msgID)

		// wrote headers
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Header().Add("Status", "404")
		w.Header().Add("Server", "Trinket")
		w.Header().Add("Connection", "close")
		w.Header().Add("Cache-Control", "private, max-age=0, must-revalidate")

		// write 404
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Message not processed"))

		return
	}

	// parse user and password pair then handle invalid auth
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Printf("auth try: %s %s\n", username, password)
	if username == "" || password == "" || !checkAuth(username, password) {
		log.Printf("request failed, invalid auth, %s : %s\n",
			username, password)

		// wrote headers
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.Header().Add("Status", "503")
		w.Header().Add("Server", "Trinket")
		w.Header().Add("Connection", "close")
		w.Header().Add("Cache-Control", "private, max-age=0, must-revalidate")

		// write 503
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Invalid username/password"))

		return
	}

	// all OK, return request response

	// handle headers
	w.Header().Add("Content-Type", "application/xml; charset=utf-8")
	w.Header().Add("Status", "200")
	w.Header().Add("Server", "Trinket")
	w.Header().Add("Connection", "close")
	w.Header().Add("Cache-Control", "private, max-age=0, must-revalidate")

	// write response
	w.WriteHeader(http.StatusOK)
	w.Write(listingQueue[msgID-1].Response)

	log.Println("served GET")
}
