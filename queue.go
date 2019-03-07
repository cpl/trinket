package main

import (
	"log"
	"time"
)

// Listing is an enqueued request which can be processed or pending, after the
// request is processed the client can retrieve the request response from the
// given URI (which contains the listing ID).
type Listing struct {
	Status string

	username string
	password string

	Request  []byte
	Response []byte
}

var listingQueue []*Listing

// AppendToQueue takes the request. Return the message and returns the ID
// which is then processed into the URI.
func AppendToQueue(request []byte) int {
	// create listing with request
	newListing := new(Listing)
	newListing.Request = request
	newListing.Status = "pending"

	// append to queue
	listingQueue = append(listingQueue, newListing)

	// return id
	return len(listingQueue)
}

// ProcessListings is a background agent that runs on it's own thread,
// processing requests from the listing queue. After each listing the agent will
// go to sleep for the given duration.
// lastID acts as the last processed listing. When started it is 0, and always
// increments by one after processing a listing. This way lastID can be used
// as the Index for the current listing to process. Listings IDs are in
// range[1:] while the listingQueue is [0:].
func ProcessListings(delay time.Duration) {
	log.Println("started processing listings")
	lastID := 0
	for {
		// check for new listings
		if lastID < len(listingQueue) {
			log.Printf("processing %d\n", lastID+1)

			// get current listing to process and send it to parser
			parseListing(listingQueue[lastID])

			// next
			lastID++
		}

		// wait before processing next listing
		time.Sleep(delay)
	}
}
