import json
import requests

r = requests.get('http://localhost:8080/events', stream=True)

for line in r.iter_lines():
    if line:
        decoded_line = line.decode('utf-8')
        print(decoded_line)
