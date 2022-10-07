// Package go-cities.json is Go's mirror for https://github.com/lutangar/cities.json with Go's embed support.
package gocitiesjson

import (
	_ "embed"
	"encoding/json"
)

//go:embed upstream/cities.json
var citiesbytes []byte

type City struct {
	Country string `json:"country"`
	Name    string `json:"name"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
}

var Cities []*City

func init() {
	if err := json.Unmarshal(citiesbytes, &Cities); err != nil {
		panic(err)
	}
}
