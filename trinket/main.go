package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var slots []string
var username string
var password string

func init() {
	// extract slot count from argv
	slotsCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	// generate empty slots
	slots = make([]string, slotsCount)
	for idx := range slots {
		slots[idx] = "free"
	}

	log.Printf("generated %d empty slots\n", slotsCount)

	// user and password parsing
	username = os.Args[3]
	password = os.Args[4]
	log.Printf("assigned user '%s' with password '%s'\n",
		username, password)

}

func main() {
	// create router
	r := mux.NewRouter()

	// define handlers and routes
	r.HandleFunc("/queue/enqueue", handleEnqueue).Methods(http.MethodPut)
	r.HandleFunc("/queue/msg/{msg_id}", handleEnqueue).Methods(http.MethodGet)

	// user views
	r.HandleFunc("/booking", handleBookings).Methods(http.MethodGet)
	r.HandleFunc("/booking/available", handleAvailable).Methods(http.MethodGet)
	r.HandleFunc("/queue", handleQueue).Methods(http.MethodGet)

	// create server
	srv := &http.Server{
		Handler:      context.ClearHandler(r),
		Addr:         "localhost:" + os.Args[1],
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	log.Println("starting server on localhost:" + os.Args[1])

	// start server
	log.Fatal(srv.ListenAndServe())
}
