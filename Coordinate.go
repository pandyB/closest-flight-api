package main

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func newCoordinateFromLatLong(Latitude float64, Longitude float64) *Coordinate {
	return &Coordinate{Latitude, Longitude}
}
