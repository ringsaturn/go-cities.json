from typing import List

class City(object):
    lng: float
    lat: float
    country: str
    name: str


def random_city() -> City: ...

def all_cities() -> List[City]: ...
