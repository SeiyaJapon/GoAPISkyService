package flight

type FlightDataProvider interface {
	// Fetch all flights in raw form
	FetchFlights() ([]Flight, error)
}
