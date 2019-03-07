package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleEnqueuePUT(w http.ResponseWriter, r *http.Request) {
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
	AppendToQueue(bodyData)

	// append request to queue and format URI with msg id
	uri := fmt.Sprintf(
		"<msg_uri>http://localhost:%d/queue/msg/%d</msg_uri>",
		port, AppendToQueue(bodyData))

	// handle headers
	w.Header().Add("Content-Type", "application/xml; charset=utf-8")
	w.Header().Add("Status", "200")
	w.Header().Add("Server", "Trinket")
	w.Header().Add("Connection", "close")
	w.Header().Add("Cache-Control", "private, max-age=0, must-revalidate")

	// write response URI
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(uri))
}

func handleEnqueueGET(w http.ResponseWriter, r *http.Request) {

}
