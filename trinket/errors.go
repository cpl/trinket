package main

import "encoding/xml"

// ResponseError contains just an error code and the error string as defined
// by the Exercise 2 labscript.
type ResponseError struct {
	XMLName xml.Name
	Code    int    `xml:"code"`
	Body    string `xml:"body"`
}

var responseErrorAuth = ResponseError{
	XMLName: xml.Name{Local: "response"},
	Code:    401,
	Body:    "Action failed due to an invalid username password",
}

var responseErrorInvalid = ResponseError{
	XMLName: xml.Name{Local: "response"},
	Code:    510,
	Body:    "Invalid Request",
}
