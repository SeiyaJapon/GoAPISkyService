package service

import (
	flight2 "GoAPISkyService/internal/flight_context/domain/flight"
)

type FetchAllFlightsService struct {
	flightInterface flight2.FlightRepositoryInterface
}

func NewFetchAllFlightsService(flightInterface flight2.FlightRepositoryInterface) *FetchAllFlightsService {
	return &FetchAllFlightsService{
		flightInterface: flightInterface,
	}
}

func (s *FetchAllFlightsService) FetchAllFlightsService() ([]flight2.Flight, error) {
	return s.flightInterface.FetchAllFlights()
}
