package bootstrap

import (
	"GoAPISkyService/internal/flight_context/infrastructure/bootstrap/bus"
	commandInterface "GoAPISkyService/internal/shared_context/cqrs/application/command"
	"GoAPISkyService/internal/shared_context/cqrs/application/query"
	"reflect"
)

func RegisterCommands(commandBus *bus.InMemoryCommandBus, commandHandlers ...interface{}) {
	for _, handler := range commandHandlers {
		handlerType := reflect.TypeOf(handler)
		if handlerType.Kind() != reflect.Ptr || handlerType.Elem().Kind() != reflect.Struct {
			panic("Command handler must be a pointer to a struct")
		}

		handlerMethod, ok := handlerType.Elem().MethodByName("Handle")
		if !ok || handlerMethod.Type.NumIn() != 2 || handlerMethod.Type.NumOut() != 0 {
			panic("Command handler must have a Handle method with one argument and no return values")
		}

		commandType := handlerMethod.Type.In(1)

		commandPointer := reflect.New(commandType.Elem()).Interface()
		if _, ok := commandPointer.(commandInterface.CommandInterface); !ok {
			panic("Command handler must handle a type that implements CommandInterface")
		}

		commandBus.RegisterHandler(commandType, handler.(commandInterface.CommandHandlerInterface))
	}
}

func RegisterQueries(queryBus *bus.InMemoryQueryBus, queryHandlers ...interface{}) {
	for _, handler := range queryHandlers {
		handlerType := reflect.TypeOf(handler)
		if handlerType.Kind() != reflect.Ptr || handlerType.Elem().Kind() != reflect.Struct {
			panic("Query handler must be a pointer to a struct")
		}

		handlerMethod, ok := handlerType.Elem().MethodByName("Handle")
		if !ok || handlerMethod.Type.NumIn() != 2 || handlerMethod.Type.NumOut() != 2 {
			panic("Query handler must have a Handle method with one argument and two return values")
		}

		queryType := handlerMethod.Type.In(1)

		queryPointer := reflect.New(queryType.Elem()).Interface()
		if _, ok := queryPointer.(query.QueryInterface); !ok {
			panic("Query handler must handle a type that implements QueryInterface")
		}

		queryBus.RegisterHandler(queryType, handler.(query.QueryHandlerInterface))
	}
}
