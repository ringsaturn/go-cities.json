package gocitiesjson_test

import (
	"testing"

	gocitiesjson "github.com/ringsaturn/go-cities.json"
)

func TestInitFine(t *testing.T) {
	if len(gocitiesjson.Cities) < 10000 {
		t.Fatalf("too less")
	}
}
