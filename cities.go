// Package go-cities.json is Go's mirror for https://github.com/lutangar/cities.json with Go's embed support.
package gocitiesjson

import (
	_ "embed"
	"encoding/json"
	"strconv"
)

//go:embed upstream/cities.json
var citiesbytes []byte

type City struct {
	Country string  `json:"country"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Admin1  string  `json:"admin1"`
	Admin2  string  `json:"admin2"`
}

var Cities []*City

func init() {
	type RawCity struct {
		Country string `json:"country"`
		Name    string `json:"name"`
		Lat     string `json:"lat"`
		Lng     string `json:"lng"`
		Admin1  string `json:"admin1"`
		Admin2  string `json:"admin2"`
	}

	var rawcities []*RawCity
	if err := json.Unmarshal(citiesbytes, &rawcities); err != nil {
		panic(err)
	}
	for _, rawcity := range rawcities {
		lat, err := strconv.ParseFloat(rawcity.Lat, 64)
		if err != nil {
			panic(err)
		}
		lng, err := strconv.ParseFloat(rawcity.Lng, 64)
		if err != nil {
			panic(err)
		}
		Cities = append(Cities, &City{
			Lat:     lat,
			Lng:     lng,
			Country: rawcity.Country,
			Name:    rawcity.Name,
			Admin1:  rawcity.Admin1,
			Admin2:  rawcity.Admin2,
		})
	}
}
