package main

import "encoding/xml"

const REQUEST_RESERVE = "reserve"
const REQUEST_CANCEL = "cancel"
const REQUEST_AVAILABILITY = "availability"
const REQUEST_BOOKINGS = "bookings"

type requestBase struct {
	XMLName xml.Name

	ID       int64  `xml:"request_id"`
	Username string `xml:"username"`
	Password string `xml:"password"`
}

/*
RequestReserve is the expected XML file representation of
a reserve request.
This comes in as attached to reserve requests of type PUT.

XML FORM:
<reserve>
  <request_id>id</request_id>
  <username>username</username>
  <password>password</password>
  <slot_id>slot_id </slot_id>
</reserve>
*/
type RequestReserve struct {
	XMLName xml.Name

	ID       int64  `xml:"request_id"`
	Username string `xml:"username"`
	Password string `xml:"password"`

	SlotID int `xml:"slot_id"`
}

/*
RequestAvailability is the expected XML file representation of
a availability request.
This comes in as attached to availability requests of type PUT.

XML FORM:
<availability>
  <request_id>id</request_id>
  <username>username</username>
  <password>password</password>
</availability>
*/
type RequestAvailability struct {
	XMLName xml.Name

	ID       int64  `xml:"request_id"`
	Username string `xml:"username"`
	Password string `xml:"password"`
}

/*
RequestCancel is the expected XML file representation of
a cancel request.
This comes in as attached to cancel requests of type PUT.

XML FORM:
<cancel>
  <request_id>id</request_id>
  <username>username</username>
  <password>password</password>
  <slot_id>slot_id </slot_id>
</cancel>
*/
type RequestCancel struct {
	XMLName xml.Name

	ID       int64  `xml:"request_id"`
	Username string `xml:"username"`
	Password string `xml:"password"`

	SlotID int `xml:"slot_id"`
}

/*
RequestBookings is the expected XML file representation of
a bookings request.
This comes in as attached to bookings requests of type PUT.

XML FORM:
<bookings>
  <request_id>id</request_id>
  <username>username</username>
  <password>password</password>
</bookings>
*/
type RequestBookings struct {
	XMLName xml.Name

	ID       int64  `xml:"request_id"`
	Username string `xml:"username"`
	Password string `xml:"password"`
}
