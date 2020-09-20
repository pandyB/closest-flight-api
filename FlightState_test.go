package main

import (
	"testing"
)

func TestFlightState(t *testing.T) {

	var expectedTxnAddress string = "a2ebbd"
	var expectedCallsign string = "ASA404  "
	var expectedOriginCountry string = "United States"
	var expectedTimePosition float64 = 1597440639
	var expectedLastContact float64 = 1597440639
	var expectedLongitude float64 = -93.2459
	var expectedLatitude float64 = 38.1067

	testSequence := []interface{}{expectedTxnAddress,
		expectedCallsign,
		expectedOriginCountry,
		expectedTimePosition,
		expectedLastContact,
		expectedLongitude,
		expectedLatitude, 11277.6, false, 250.69, 108.55, 0.0, nil, 11826.24, "0036", false, 0}
	fs := NewFlightStateFromSequence(testSequence)

	if *fs.TxnAddress != expectedTxnAddress {
		t.Errorf("Expected: %s, got: %s\n", expectedTxnAddress, *fs.TxnAddress)
	}
	if *fs.Callsign != expectedCallsign {
		t.Errorf("Expected: %s, got: %s\n", expectedCallsign, *fs.Callsign)
	}
	if *fs.OriginCountry != expectedOriginCountry {
		t.Errorf("Expected: %s, got: %s\n", expectedOriginCountry, *fs.Callsign)
	}
	if *fs.TimePosition != expectedTimePosition {
		t.Errorf("Expected: %f, got: %f\n", expectedTimePosition, *fs.TimePosition)
	}
	if *fs.LastContact != expectedLastContact {
		t.Errorf("Expected: %f, got: %f\n", expectedLastContact, *fs.LastContact)
	}
	if *fs.Longitude != expectedLongitude {
		t.Errorf("Expected: %f, got: %f\n", expectedLongitude, *fs.Longitude)
	}
	if *fs.Latitude != expectedLatitude {
		t.Errorf("Expected: %f, got: %f\n", expectedLatitude, *fs.Latitude)
	}
}
