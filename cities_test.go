package gocitiesjson_test

import (
	"encoding/json"
	"os"
	"testing"

	gocitiesjson "github.com/ringsaturn/go-cities.json"
)

func TestInitFine(t *testing.T) {
	if len(gocitiesjson.Cities) < 10000 {
		t.Fatalf("too less")
	}
}

type FeatureCollection struct {
	Type     string     `json:"type"` // FeatureCollection
	Features []Features `json:"features"`
}

type Features struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Geometry   Geometry               `json:"geometry"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func TestGenGeoJSON(t *testing.T) {
	featureCll := &FeatureCollection{
		Type:     "FeatureCollection",
		Features: make([]Features, 0),
	}
	for _, city := range gocitiesjson.Cities {
		featureCll.Features = append(featureCll.Features, Features{
			Type: "Feature",
			Properties: map[string]interface{}{
				"Country": city.Country,
				"Name":    city.Name,
				"Admin1":  city.Admin1,
				"Admin2":  city.Admin2,
			},
			Geometry: Geometry{
				Type:        "Point",
				Coordinates: []float64{city.Lng, city.Lat},
			},
		})
	}
	featureBytes, err := json.Marshal(featureCll)
	if err != nil {
		panic(err)
	}
	_ = os.WriteFile("gocities.geojson", featureBytes, 0644)
}
