package haversine

import (
	"math"
)

const (
	m  = 6371000 // radius of the earth in meters
	km = 6371    // radius of the earth in km
	mi = 3958    // radius of the earth in miles
)

// Coordinate ..
type Coordinate struct {
	Lat float64
	Lng float64
}

func degreeToRadians(x float64) float64 {
	return x * math.Pi / 180
}

// Calculate ..
func Calculate(a, b Coordinate) (meters, kilometers, miles float64) {
	// change coordinates to radians
	lat1 := degreeToRadians(a.Lat)
	lat2 := degreeToRadians(b.Lat)
	deltaLat := degreeToRadians(b.Lat - a.Lat)
	deltaLng := degreeToRadians(b.Lng - a.Lng)

	// calculate haversine formula
	x := math.Pow(math.Sin(deltaLat/2), 2) + (math.Cos(lat1) * math.Cos(lat2) * math.Pow(math.Sin(deltaLng/2), 2))
	y := 2 * math.Atan2(math.Sqrt(x), math.Sqrt(1-x))
	meters = y * m
	kilometers = y * km
	miles = y * mi
	return meters, kilometers, miles
}
