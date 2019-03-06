package main

import "time"

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

	// fill in listing
	newListing := new(Listing)

	newListing.Request = request

	// append to queue
	listingQueue = append(listingQueue, newListing)

	return len(listingQueue) - 1
}

// ProcessListings is a background agent that runs on it's own thread,
// processing requests from the listing queue. After each listing the agent will
// go to sleep for the given duration.
func ProcessListings(delay time.Duration) {
	lastID := -1
	for {
		if lastID < len(listingQueue) {

		}

		time.Sleep(delay)
	}
}
