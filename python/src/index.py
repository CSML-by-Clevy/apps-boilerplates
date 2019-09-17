import requests


def handler(event, context):
    url = ('https://randomfox.ca/floof/')
    response = requests.get(url)
    resjson = response.json()
    return resjson
