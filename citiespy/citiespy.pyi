from typing import List

class City(object):
    lng: float
    lat: float
    country: str
    name: str
    admin1: str
    admin2: str


def random_city() -> City: ...

def all_cities() -> List[City]: ...
