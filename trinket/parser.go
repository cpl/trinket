package main

import (
	"encoding/xml"
	"log"
)

func parseListing(listing *Listing) {

	// default is no username or password, because the request may be invalid
	listing.username = ""
	listing.password = ""

	// TODO check request ID is unique

	// validate request
	// this means checking it has a username, password and unique ID
	var baseRequest requestBase
	if err := xml.Unmarshal(listing.request, &baseRequest); err != nil {
		log.Printf("invalid request format")
		listing.response, _ = xml.Marshal(responseErrorInvalid)
		return
	}

	// if format is valid, assign username and password to listing
	// they will later be checked against the server users for the
	// proper request types
	listing.username = baseRequest.Username
	listing.password = baseRequest.Password

	// ! this should happen only for some requests Reserve, Cancel, Bookings
	// validate username and password
	// if !checkAuth(baseRequest.Username, baseRequest.Password) {
	// 	log.Printf("invalid auth for request %d\n", baseRequest.ID)
	// 	ret, _ := xml.Marshal(responseErrorAuth)
	// 	return
	// }

	// validate and handle request
	switch baseRequest.XMLName.Local {
	// --------------------------------------------------------------------
	case REQUEST_AVAILABILITY:
		var req RequestAvailability

		// fully parse request
		if err := xml.Unmarshal(listing.request, &req); err != nil {
			log.Printf("invalid request format")
			listing.response, _ = xml.Marshal(responseErrorInvalid)
			return
		}

		// prepare response
		var res ResponseAvailability
		res.XMLName = xml.Name{Local: "response"}
		res.Code = 200

		// get all free slots
		for idx := range slots {
			if slots[idx] == "free" {
				res.Slots = append(res.Slots, idx)
			}
		}

		// create response payload
		resBinary, err := xml.Marshal(res)
		if err != nil {
			panic(err)
		}

		// assign response to listing
		listing.response = resBinary

		log.Printf("parsed request %d of type '%s'\n",
			req.ID, req.XMLName)

		return
	// --------------------------------------------------------------------
	case REQUEST_BOOKINGS:
		// var req RequestAvailability
		break
	// --------------------------------------------------------------------
	case REQUEST_CANCEL:
		// var req RequestAvailability
		break
	// --------------------------------------------------------------------
	case REQUEST_RESERVE:
		// var req RequestAvailability
		break
	// --------------------------------------------------------------------
	default:
		panic("should not reach this")
	}
}
