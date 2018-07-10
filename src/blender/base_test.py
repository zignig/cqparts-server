import json
import requests

r = requests.get('http://localhost:8080/render', stream=True)

def fetch(url):
    pass

for line in r.iter_lines():
    if line:
        decoded_line = line.decode('utf-8')
        print(decoded_line)
