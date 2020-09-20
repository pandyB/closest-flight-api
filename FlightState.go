package main

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
	TxnAddress    *string
	Callsign      *string
	OriginCountry *string
	TimePosition  *float64
	LastContact   *float64
	Longitude     *float64
	Latitude      *float64
}

func NewFlightStateFromSequence(dataSeq []interface{}) *FlightState {
	txnAddress, _ := dataSeq[0].(string)
	callsign, _ := dataSeq[1].(string)
	originCountry, _ := dataSeq[2].(string)
	timePosition, _ := dataSeq[3].(float64)
	lastContact, _ := dataSeq[4].(float64)
	longitude, _ := dataSeq[5].(float64)
	latitude, _ := dataSeq[6].(float64)

	return &FlightState{
		TxnAddress:    &txnAddress,
		Callsign:      &callsign,
		OriginCountry: &originCountry,
		TimePosition:  &timePosition,
		LastContact:   &lastContact,
		Longitude:     &longitude,
		Latitude:      &latitude}
}
