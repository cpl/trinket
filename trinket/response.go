package main

import "encoding/xml"

type responseBase struct {
	xmlName xml.Name

	Code int    `xml:"code"`
	Body string `xml:"body"`
}

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
	responseBase

	Slots []int `xml:"body>availability>slot_id"`
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
	responseBase

	Slots []int `xml:"body>bookings>slot_id"`
}
