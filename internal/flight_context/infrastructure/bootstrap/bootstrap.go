package bootstrap

import (
	"GoAPISkyService/internal/flight_context/infrastructure/bootstrap/cqrs"
	"GoAPISkyService/internal/flight_context/infrastructure/bootstrap/cqrs/bus"
	"GoAPISkyService/internal/flight_context/infrastructure/bootstrap/db"
	"database/sql"
	"reflect"
)

type AppDependencies struct {
	DB         *sql.DB
	CommandBus *bus.InMemoryCommandBus
	QueryBus   *bus.InMemoryQueryBus
}

func InitDependencies() (*AppDependencies, error) {
	dbConfig, err := db.InitDB()
	if err != nil {
		return nil, err
	}

	commandBus := bus.NewInMemoryCommandBus()
	queryBus := bus.NewInMemoryQueryBus()

	for _, registration := range cqrs.CommandHandlers {
		commandBus.RegisterHandler(
			reflect.TypeOf(registration.Command),
			registration.Handler,
		)
	}

	for _, registration := range cqrs.QueryHandlers {
		queryBus.RegisterHandler(
			reflect.TypeOf(registration.Query),
			registration.Handler,
		)
	}

	return &AppDependencies{
		DB:         dbConfig,
		CommandBus: commandBus,
		QueryBus:   queryBus,
	}, nil
}
