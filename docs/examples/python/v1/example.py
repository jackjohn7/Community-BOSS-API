import requests

# make the request
r = requests.get("https://community-boss-api-development.up.railway.app/beta/quarters")

# parse the response body into JSON (r.status_code should be 200)
print(r.json())

# going further, you could parse the json into pydantic classes or dataclasses for
#  additional typesafety
