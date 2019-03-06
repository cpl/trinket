import requests

URL = "http://localhost:3000/queue/enqueue"
# URL = "<REDACTED>"

headers = {
    "Content-Type": "application/xml",
    "Accept": "application/xml"
}

file_data = None
with open('req.xml', 'rb') as fp:
    file_data = fp.read()


response = requests.put(URL, data=file_data, headers=headers)

print(response.status_code, response.content)



# BEHAVIOR NOTES
# if message is pending: 404 Message unavailable
# if wrong auth for Availability -> OK ???

