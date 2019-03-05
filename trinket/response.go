package main

/*
ResponseBase is the expected response minimal format.
The body field may contain extra information such as:
strings or more XML which will further be parsed into
other Response... type structs.

XML FORM:
<response>
  <code>code</code>
  <body> ... </body>
</response>
*/
type ResponseBase struct {
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
    <slot_id> 15 </slot_id>
  </availability>
</body> </response>
*/
type ResponseAvailability struct {
	Code  int   `xml:"code"`
	Slots []int `xml:"body>availability>slot_id"`
}
