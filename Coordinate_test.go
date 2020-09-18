package main

import (
	"fmt"
	"testing"
)

func TestCoordinate(t *testing.T) {
	myCoor := newCoordinateFromLatLong(12.34, -100.9)

	fmt.Println(myCoor.Latitude)
	fmt.Println(myCoor.Longitude)
}
