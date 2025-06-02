package repositories

import (
	"GoAPISkyService/internal/flight_context/domain/flight"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type OpenskyFetchAllFlights struct {
	apiURL string
}

func NewOpenskyFetchAllFlights(apiURL string) *OpenskyFetchAllFlights {
	return &OpenskyFetchAllFlights{
		apiURL: apiURL,
	}
}

func (o *OpenskyFetchAllFlights) FetchAllFlights() ([]flight.Flight, error) {
	resp, err := http.Get(o.apiURL)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error fetching flights: " + resp.Status)
	}

	var flights []flight.Flight
	if err := json.NewDecoder(resp.Body).Decode(&flights); err != nil {
		return nil, err
	}

	return flights, nil
}
