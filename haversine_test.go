package haversine

import (
	"log"
	"testing"
)

// TestHaversine ..
func TestHaversine(t *testing.T) {
	coordinatA := Coordinate{
		Lat: -6.2426929,
		Lng: 106.8545669,
	}
	coordinatB := Coordinate{
		Lat: -6.2427037,
		Lng: 106.8547062,
	}
	m, km, mi := Calculate(coordinatA, coordinatB)
	log.Println(m)
	log.Println(km)
	log.Println(mi)
}
