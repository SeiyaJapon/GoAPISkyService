package bus

import (
	"GoAPISkyService/internal/shared_context/cqrs/application/exceptions"
	"GoAPISkyService/internal/shared_context/cqrs/application/query"
	"reflect"
)

type InMemoryQueryBus struct {
	handlers map[reflect.Type]query.QueryHandlerInterface
}

func NewInMemoryQueryBus() *InMemoryQueryBus {
	return &InMemoryQueryBus{
		handlers: make(map[reflect.Type]query.QueryHandlerInterface),
	}
}

func (bus *InMemoryQueryBus) RegisterHandler(query query.QueryInterface, handler query.QueryHandlerInterface) {
	queryType := reflect.TypeOf(query)
	if _, exists := bus.handlers[queryType]; exists {
		panic("Query handler already registered for query type: " + queryType.String())
	}
	bus.handlers[queryType] = handler
}

func (bus *InMemoryQueryBus) Ask(query query.QueryInterface) (query.QueryResultInterface, error) {
	queryType := reflect.TypeOf(query)
	handler, exists := bus.handlers[queryType]
	if !exists {
		return nil, exceptions.ErrQueryHandlerNotFound
	}
	return handler.Handle(query)
}
