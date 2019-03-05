package main

import "encoding/xml"

type requestBase struct {
	xmlNodeName xml.Name

	ID       int    `xml:"request_id"`
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
	requestBase
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
	requestBase
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
	requestBase
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
	requestBase
}
