from unittest import TestCase

from citiespy import random_city


class TestCities(TestCase):
    def test_random(self):
        _ = random_city()
