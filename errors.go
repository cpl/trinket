package main

import "encoding/xml"

// ResponseError contains just an error code and the error string as defined
// by the Exercise 2 labscript.
type ResponseError struct {
	XMLName xml.Name
	Code    int    `xml:"code"`
	Body    string `xml:"body"`
}

// ResponseErrorAuth if username password pair does not match db or does not
// match request or does not exist, return this error.
var ResponseErrorAuth = ResponseError{
	XMLName: xml.Name{Local: "response"},
	Code:    401,
	Body:    "Action failed due to an invalid username password",
}

// ResponseErrorInvalid if the format of the request was not as expected.
var ResponseErrorInvalid = ResponseError{
	XMLName: xml.Name{Local: "response"},
	Code:    510,
	Body:    "Invalid Request",
}

// ResponseErrorDuplicateID is a custom form of Invalid Error. This request
// does not exist on the main server as the labscript does not mention the
// actual behavior for this.
var ResponseErrorDuplicateID = ResponseError{
	XMLName: xml.Name{Local: "response"},
	Code:    510,
	Body:    "Invalid Request - Duplicate ID",
}
