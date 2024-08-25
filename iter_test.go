//go:build go1.23
// +build go1.23

package gocitiesjson_test

import (
	"fmt"

	gocitiesjson "github.com/ringsaturn/go-cities.json"
)

func ExampleAll() {
	c := 0
	for city := range gocitiesjson.All() {
		_ = city
		c++
	}
	fmt.Println(c > 10000)
	// Output: true
}

func ExampleFilter() {
	c := 0

	bboxFilter := func(city *gocitiesjson.City) bool {
		return city.Lng > 0 && city.Lng < 10 && city.Lat > 0 && city.Lat < 10
	}
	for city := range gocitiesjson.Filter(bboxFilter) {
		_ = city
		c++
	}

	fmt.Println(c > 0)
	// Output: true
}
