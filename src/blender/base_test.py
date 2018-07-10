import json
import requests
import time

def tryAgain(retries=0):
    if retries > 50: return
    try:
        r = requests.get('http://localhost:8080/render', stream=True)
        for line in r.iter_lines():
            if line:
                decoded_line = line.decode('utf-8')
                print(decoded_line)
    except:
        time.sleep(15)
        retries+=1
        print ("retries"+str(retries))
        tryAgain(retries)

tryAgain()
