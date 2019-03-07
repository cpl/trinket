package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

var slots []string
var port int

var usersMap map[string]string
var usersSlots map[string][]int

var maxBookedSlots int

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
	users := strings.Split(os.Args[3], " ")
	pswds := strings.Split(os.Args[4], " ")
	if len(users) == 0 || len(pswds) == 0 || len(users) != len(pswds) {
		log.Println("invalid number of users and passwords")
	} else {
		usersMap = make(map[string]string, len(users))
		usersSlots = make(map[string][]int, len(users))

		for idx := range users {
			log.Printf("adding user %s: %s\n", users[idx], pswds[idx])

			// create name password entry
			usersMap[users[idx]] = pswds[idx]

			// create list of booked slots for each user
			usersSlots[users[idx]] = make([]int, 0, maxBookedSlots)
		}
	}

	// parse number of max possible booked slots
	maxBookedSlots, err = strconv.Atoi(os.Args[5])
	if err != nil {
		panic(err)
	}
	log.Printf("maximum number of booked slots is %d\n", maxBookedSlots)

	// parse port
	port, err = strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
}

func main() {
	// create router
	r := mux.NewRouter()

	// define handlers and routes for PUT requests and GET responses
	r.HandleFunc("/queue/enqueue", handleEnqueuePUT).Methods(http.MethodPut)
	r.HandleFunc("/queue/msg/{msg_id}", handleEnqueueGET).Methods(http.MethodGet)

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

	// start processing requests
	go ProcessListings(1 * time.Second)

	// start server
	log.Println("starting server on localhost:" + os.Args[1])
	log.Fatal(srv.ListenAndServe())
}
