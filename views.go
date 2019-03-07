package main

import (
	"fmt"
	"net/http"
)

func handleQueue(w http.ResponseWriter, r *http.Request) {
	for idx, listing := range listingQueue {
		w.Write([]byte(
			fmt.Sprintf("%8d %10s %8s\n",
				idx+1, listing.Status, listing.username)))
	}
}

func handleBookings(w http.ResponseWriter, r *http.Request) {
	for idx, slot := range slots {
		w.Write([]byte(fmt.Sprintf("slot %8d : %8s\n", idx+1, slot)))
	}
}

func handleAvailable(w http.ResponseWriter, r *http.Request) {
	for idx, slot := range slots {
		if slot == "free" {
			w.Write([]byte(fmt.Sprintf("%d\n", idx+1)))
		}
	}
}
