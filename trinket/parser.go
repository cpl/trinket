package main

import (
	"encoding/xml"
	"log"
)

func parseRequest(data []byte) []byte {

	// validate request
	var baseRequest requestBase
	if err := xml.Unmarshal(data, &baseRequest); err != nil {
		log.Printf("invalid request format")
		ret, _ := xml.Marshal(responseErrorInvalid)
		return ret
	}

	// validate username and password
	if !checkAuth(baseRequest.Username, baseRequest.Password) {
		log.Printf("invalid auth for request %d\n", baseRequest.ID)
		ret, _ := xml.Marshal(responseErrorAuth)
		return ret
	}

	// validate and handle request
	switch baseRequest.XMLName.Local {
	// --------------------------------------------------------------------
	case REQUEST_AVAILABILITY:
		var req RequestAvailability

		if err := xml.Unmarshal(data, &req); err != nil {
			panic(err)
		}

		var res ResponseAvailability

		res.XMLName = xml.Name{Local: "response"}
		res.Code = 200
		for idx := range slots {
			if slots[idx] == "free" {
				res.Slots = append(res.Slots, idx)
			}
		}

		xmlResponse, err := xml.Marshal(res)
		if err != nil {
			panic(err)
		}

		log.Printf("parsed request %d of type '%s'\n",
			req.ID, req.XMLName)

		return xmlResponse
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

	return nil
}
