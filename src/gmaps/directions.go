package main

import (

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"fmt"
	"log"
)

func main() {

	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyAdX7F3EcjKq8WWpLA8D8pTSHfYC4EQRWI"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r :=  &maps.DistanceMatrixRequest{
		Origins: []string{"evesham, UK"},
		Destinations: []string{"Oxford, UK"},
		Units: maps.UnitsImperial,
		Language: "en",
		DepartureTime: "now",
	}

	route, err := c.DistanceMatrix(context.Background(),r)

	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	fmt.Println("Duration in minutes: %f\n", route.Rows[0].Elements[0].DurationInTraffic.Minutes())
}
