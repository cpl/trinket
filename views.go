package main

import (
	"fmt"
	"net/http"
)

func handleQueue(w http.ResponseWriter, r *http.Request) {
	for idx, slot := range slots {
		w.Write([]byte(fmt.Sprintf("slot %d : %s\n", idx+1, slot)))
	}
}

func handleBookings(w http.ResponseWriter, r *http.Request) {

}

func handleAvailable(w http.ResponseWriter, r *http.Request) {

}
