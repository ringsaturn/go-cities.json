from unittest import TestCase

from citiespy import all_cities, random_city


class TestCities(TestCase):
    def test_get_random(self):
        _ = random_city()

    def test_get_all(self):
        cities = all_cities()
        self.assertGreater(len(cities), 100000)
