// Package go-cities.json is Go's mirror for https://github.com/lutangar/cities.json with Go's embed support.
package gocitiesjson

type City struct {
	Country string `json:"country"`
	Name    string `json:"name"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
}
