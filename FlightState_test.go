package main

import (
	"fmt"
	"testing"
)

func TestFlightState(t *testing.T) {
	dummy_rec := []interface{}{"a2ebbd", "ASA404  ", "United States", 1597440639, 1597440639, -93.2459, 38.1067, 11277.6, false, 250.69, 108.55, 0.0, nil, 11826.24, "0036", false, 0}
	fs := FlightState{data: dummy_rec}

	fmt.Println(*fs.getTxnAddress())
	fmt.Println(*fs.getCallSign())
	fmt.Println(*fs.getSensors())
	fmt.Println(*fs.getLong())
	fmt.Println(*fs.getLat())
}
