package main

import (
	"fmt"
	"testing"
)

func TestCoordinate(t *testing.T) {
	myCoor, _ := NewCoordinateFromLatLong(12.34, -100.9)

	fmt.Println(myCoor.Latitude)
	fmt.Println(myCoor.Longitude)
}

func TestInvalidCoordinate(t *testing.T) {
	_, err1 := NewCoordinateFromLatLong(-91.56, -100.9)
	if err1 == nil {
		t.Errorf("Expected error for invalid coordinates")
	}

	_, err2 := NewCoordinateFromLatLong(98, -100.9)
	if err2 == nil {
		t.Errorf("Expected error for invalid coordinates")
	}

	_, err3 := NewCoordinateFromLatLong(12.34, -180.01)
	if err3 == nil {
		t.Errorf("Expected error for invalid coordinates")
	}

	_, err4 := NewCoordinateFromLatLong(12.34, 189.87)
	if err4 == nil {
		t.Errorf("Expected error for invalid coordinates")
	}
}

func TestHashCoordinate(t *testing.T) {
	coordA, _ := NewHashedCoordinateFromLatLong(45.1, -86.07)
	coordB, _ := NewHashedCoordinateFromLatLong(44.27, -87.88)
	fmt.Println(coordA.IsNear(coordB))
	coordC, _ := NewHashedCoordinateFromLatLong(44.27, -40.88)
	fmt.Println(coordA.IsNear(coordC))
}
