package main

import "math"

// Source: (MIT license, copyright (c) 2014 Chris Veness)
// https://www.movable-type.co.uk/scripts/latlong.html
// https://github.com/chrisveness/geodesy/
func getDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	const R = 6371e3             // meters
	phi1 := lat1 * math.Pi / 180 // φ, λ in radians
	phi2 := lat2 * math.Pi / 180
	deltaPhi := (lat2 - lat1) * math.Pi / 180
	deltaLambda := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
		math.Cos(phi1)*math.Cos(phi2)*
			math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	d := R * c // in meters
	return d
}
