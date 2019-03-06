import requests


headers = {
    "Content-Type": "application/xml",
    "Accept": "application/xml"
}

file_data = None
with open('req.xml', 'rb') as fp:
    file_data = fp.read()


response = requests.put(
    "http://localhost:3000/queue/enqueue",
    data=file_data, headers=headers)

print(response.status_code, response.content)
