package cqrs

import (
	"GoAPISkyService/internal/flight_context/infrastructure/bootstrap/cqrs/bus"
	commandInterface "GoAPISkyService/internal/shared_context/cqrs/application/command"
	"GoAPISkyService/internal/shared_context/cqrs/application/query"
	"reflect"
)

func RegisterCommands(commandBus *bus.InMemoryCommandBus, commandHandlers ...interface{}) {
	for _, handler := range commandHandlers {
		commandType := validateHandler(handler, 0, reflect.TypeOf((*commandInterface.CommandInterface)(nil)).Elem())
		commandBus.RegisterHandler(commandType, handler.(commandInterface.CommandHandlerInterface))
	}
}

func RegisterQueries(queryBus *bus.InMemoryQueryBus, queryHandlers ...interface{}) {
	for _, handler := range queryHandlers {
		queryType := validateHandler(handler, 2, reflect.TypeOf((*query.QueryInterface)(nil)).Elem())
		queryBus.RegisterHandler(queryType, handler.(query.QueryHandlerInterface))
	}
}

func validateHandler(handler interface{}, expectedReturns int, requiredInterface reflect.Type) reflect.Type {
	handlerType := reflect.TypeOf(handler)
	if handlerType.Kind() != reflect.Ptr || handlerType.Elem().Kind() != reflect.Struct {
		panic("Handler must be a pointer to a struct")
	}

	handleMethod, ok := handlerType.Elem().MethodByName("Handle")
	if !ok || handleMethod.Type.NumIn() != 2 || handleMethod.Type.NumOut() != expectedReturns {
		panic("Handler must have a Handle method with one argument and correct return values")
	}

	msgType := handleMethod.Type.In(1)
	msgPtr := reflect.New(msgType.Elem()).Interface()
	if !reflect.TypeOf(msgPtr).Implements(requiredInterface) {
		panic("Handled type does not implement the required interface")
	}

	return msgType
}
