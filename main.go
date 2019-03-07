package main

import (
	"log"
	"math/rand"
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
var usersIDs map[string][]int64

var failChance uint32

var maxBookedSlots int

func init() {
	// init rand
	// change this to a constant value if you need predictable behavior
	rand.Seed(time.Now().UnixNano())

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

	// parse number of max possible booked slots
	maxBookedSlots, err = strconv.Atoi(os.Args[5])
	if err != nil {
		panic(err)
	}
	log.Printf("maximum number of booked slots is %d\n", maxBookedSlots)

	// user and password parsing
	users := strings.Split(os.Args[3], " ")
	pswds := strings.Split(os.Args[4], " ")
	if len(users) == 0 || len(pswds) == 0 || len(users) != len(pswds) {
		log.Println("invalid number of users and passwords")
	} else {
		usersMap = make(map[string]string, len(users))
		usersSlots = make(map[string][]int, len(users))
		usersIDs = make(map[string][]int64, len(users))

		for idx := range users {
			log.Printf("adding user %s: %s\n", users[idx], pswds[idx])

			// create name password entry
			usersMap[users[idx]] = pswds[idx]

			// create list of booked slots for each user
			usersSlots[users[idx]] = make([]int, 0, maxBookedSlots)

			// create list for used IDs
			usersIDs[users[idx]] = make([]int64, 0)
		}
	}

	// parse fail chance
	fc, err := strconv.Atoi(os.Args[6])
	if err != nil {
		panic(err)
	}
	failChance = uint32(fc)
	log.Printf("fail chance is %d%%\n", failChance)

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

	// GET requests must receive username and password URL queries
	r.HandleFunc("/queue/msg/{msg_id}", handleEnqueueGET).Methods(http.MethodGet)
	r.HandleFunc(
		"/queue/msg/{msg_id}", handleEnqueueGET).Methods(
		http.MethodGet).Queries(
		"username", "{username}", "password", "{password}")

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
