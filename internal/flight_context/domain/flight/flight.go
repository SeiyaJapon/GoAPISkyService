package flight

import "time"

type Flight struct {
	ICA024       string    // Unique flight ID
	CallSign     string    // Flight call sign
	Origin       string    // Origin airport (ICAO or IATA code)
	Destination  string    // Destination airport
	Latitude     float64   // Current latitude
	Longitude    float64   // Current longitude
	Altitude     float64   // Current altitude
	AltitudeUnit string    // "feet" or "meters"
	Speed        float64   // Current speed
	SpeedUnit    string    // "knots" or "kmh"
	Heading      int       // Heading in degrees
	LastSeen     time.Time // Last observed timestamp
}
