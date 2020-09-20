package main

import (
	"fmt"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

var minLatitude float64 = -90
var maxLatitude float64 = 90
var minLongitude float64 = -180
var maxLongitude float64 = 180

func NewCoordinateFromLatLong(Latitude float64, Longitude float64) (*Coordinate, error) {
	if Latitude < minLatitude || maxLatitude < Latitude {
		err := fmt.Errorf("Latitude (%f) out of bounds [%f, %f].", Latitude, minLatitude, maxLatitude)
		return &Coordinate{0, 0}, err
	}
	if Longitude < minLongitude || maxLongitude < Longitude {
		err := fmt.Errorf("Longitude (%f) out of bounds [%f, %f].", Longitude, minLongitude, maxLongitude)
		return &Coordinate{0, 0}, err
	}
	return &Coordinate{Latitude, Longitude}, nil
}

type HashedCoordinate struct {
	Coord  *Coordinate
	hashes []string
}

func NewHashedCoordinateFromLatLong(Latitude float64, Longitude float64) (*HashedCoordinate, error) {
	coord, err := NewCoordinateFromLatLong(Latitude, Longitude)
	return &HashedCoordinate{coord, generateLocationHashes(coord)}, err
}

func generateLocationHashes(c *Coordinate) []string {
	truncLat := int(c.Latitude)
	truncLong := int(c.Longitude)
	neighborLats := []int{truncLat + 1, truncLat, truncLat - 1}
	neighborLongs := []int{truncLong + 1, truncLong, truncLong - 1}
	return generateCombinations(neighborLats, neighborLongs)
}

func generateCombinations(listA []int, listB []int) []string {
	var hashes []string
	for itemA := range listA {
		for itemB := range listB {
			hashes = append(hashes, string(itemA)+"_"+string(itemB))
		}
	}
	return hashes
}

func (hashCoord1 *HashedCoordinate) IsNear(hashCoord2 *HashedCoordinate) bool {
	return false
}
