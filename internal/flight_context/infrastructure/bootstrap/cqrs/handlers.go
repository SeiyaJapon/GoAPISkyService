package cqrs

import (
	"GoAPISkyService/internal/shared_context/cqrs/application/command"
	"GoAPISkyService/internal/shared_context/cqrs/application/query"
)

type CommandRegistration struct {
	Command command.CommandInterface
	Handler command.CommandHandlerInterface
}

type QueryRegistration struct {
	Query   query.QueryInterface
	Handler query.QueryHandlerInterface
}

var CommandHandlers = []CommandRegistration{
	// {Command: myCommand{}, Handler: &myCommandHandler{}},
}

var QueryHandlers = []QueryRegistration{
	//{
	//	Query:   fetch_all_flights.FetchAllFlightsQuery{},
	//	Handler: &fetch_all_flights.FetchAllFlightsQueryHandler{},
	//},
}
