package flight

type FlightRepositoryInterface interface {
	FetchAllFlights() ([]Flight, error)
}
