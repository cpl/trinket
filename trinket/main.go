package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {
	// create router
	r := mux.NewRouter()

	// define handlers and routes
	r.HandleFunc("/queue/enqueue", handleEnqueue).Methods(http.MethodPut)
	r.HandleFunc("/queue/msg/{msg_id}", handleEnqueue).Methods(http.MethodGet)
	r.HandleFunc("/booking", handleBookings)
	r.HandleFunc("/booking/available", handleAvailable)

	// create server
	srv := &http.Server{
		Handler:      context.ClearHandler(r),
		Addr:         "localhost:3000",
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	// start server
	log.Fatal(srv.ListenAndServe())
}
