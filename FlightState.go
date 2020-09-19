package main

import (
	"encoding/json"
	"fmt"
)

//dummy_rec := "[a2ebbd ASA404   United States 1.597440639e+09 1.597440639e+09 -93.2459 38.1067 11277.6 false 250.69 108.55 0 <nil> 11826.24 0036 false 0]"
//dummy_rec := []interface{}{"a2ebbd", "ASA404  ", "United States", 1597440639, 1597440639, -93.2459, 38.1067, 11277.6, false, 250.69, 108.55, 0.0, <nil>, 11826.24, "0036", false, 0}
// TxnAddress     *string
// Callsign       *string
// OriginCountry  *string
// TimePosition   *int
// LastContact    *int
// Longitude      *float32
// Latitude       *float32
// BaroAltitude   *float32
// OnGround       *bool
// Velocity       *float32
// TrueTrack      *float32
// VerticalRate   *float32
// Sensors        *[]int
// GeoAltitude    *float32
// Squawk         *string
// Spi            *bool
// PositionSource *int
type FlightState struct {
	data []interface{}
}

func newFlightStateFromSeqenuence(dataSeq []interface{}) *FlightState {
	newFlightState := FlightState{data: dataSeq}
	return &newFlightState
}

func (fs FlightState) getTxnAddress() *string {
	txnAddress, _ := fs.data[0].(string)
	return &txnAddress
}

func (fs FlightState) getCallSign() *string {
	callSign, _ := fs.data[1].(string)
	return &callSign
}

func (fs FlightState) getOriginCountry() *string {
	originCountry, _ := fs.data[2].(string)
	return &originCountry
}

func (fs FlightState) getLong() *float64 {
	long, _ := fs.data[5].(float64)
	return &long
}

func (fs FlightState) getLat() *float64 {
	lat, _ := fs.data[6].(float64)
	return &lat
}

func (fs FlightState) getSensors() *[]int {
	sensors, _ := fs.data[12].([]int)
	return &sensors
}

func (fs FlightState) String() string {
	s := fmt.Sprintf("{Call Sign: %s, Lat: %f,Long: %f}", *fs.getCallSign(), *fs.getLat(), *fs.getLong())
	return s
}

func (fs FlightState) MarshalJSON() ([]byte, error) {
	fmt.Println(fs)
	return json.Marshal(fs.String())
}
