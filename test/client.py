import sys
import requests


XML = """
<{req}>
    <request_id>{id}</request_id>
    <username>{user}</username>
    <password>{pasw}</password>
    <slot_id>5</slot_id>
</{req}>
"""


URL = "http://localhost:3000/queue/enqueue"
# URL = "<READACTED>"

headers = {
    "Content-Type": "application/xml",
    "Accept": "application/xml"
}

file_data = XML.format(req=sys.argv[1], id=int(sys.argv[2]), user=sys.argv[3], pasw=sys.argv[4])


response = requests.put(URL, data=file_data, headers=headers)

print(response.status_code, response.content, response.headers)



# BEHAVIOR NOTES
# if message is pending: 404 Message unavailable
# if wrong auth for Availability -> OK ???

