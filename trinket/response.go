package main

import "encoding/xml"

/*
ResponseAvailability is the response from an availability request,
if OK it will contain a 200 code and the list of available slots.

XML FORM:
<response>
<code>200</code>
<body>
  <availability>
    <slot_id> 4  </slot_id>
    <slot_id> 12 </slot_id>
    [...]
    <slot_id> 15 </slot_id>
  </availability>
</body>
</response>
*/
type ResponseAvailability struct {
	XMLName xml.Name
	Code    int   `xml:"code"`
	Slots   []int `xml:"body>availability>slot_id"`
}

/*
ResponseBookings is the response from an bookings request,
if OK it will contain a 200 code and the list of available slots.

XML FORM:
<response>
<code>200</code>
<body>
  <bookings>
    <slot_id> 4  </slot_id>
    <slot_id> 12 </slot_id>
  </bookings>
</body>
</response>
*/
type ResponseBookings struct {
	XMLName xml.Name
	Code    int   `xml:"code"`
	Slots   []int `xml:"body>bookings>slot_id"`
}

/*
ResponseReserve is the response from a reserve request,
if OK it will contain a 200 code and the confirmation for that slot.

XML FORM:
<response>
<code>200</code>
<body>
  <reserve>
    Reserved slot N
  </reserve>
</body>
</response>
*/
type ResponseReserve struct {
	XMLName xml.Name
	Code    int    `xml:"code"`
	Reserve string `xml:"body>reserve"`
}
