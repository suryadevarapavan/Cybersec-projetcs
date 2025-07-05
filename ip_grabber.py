import requests
import json
ip = str(input("Enter IP Address:"))
url= "https://api.ipbase.com/v2/info?"

key = "YOUR_API_KEY"

headers = {
     'Authorization': key,
     'ip':ip
}
response = requests.get(url, headers=headers)

if response.status_code == 200:
    data = response.json()
    print(data)
    with open("data.json","w") as f:
        json.dump(data,f ,indent=4)
else:
    print(f"Error: {response.status_code}")
