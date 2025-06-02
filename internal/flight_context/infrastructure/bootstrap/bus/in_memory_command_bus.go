package bus

import (
	"GoAPISkyService/internal/shared_context/cqrs/application/command"
	"GoAPISkyService/internal/shared_context/cqrs/application/exceptions"
	"reflect"
)

type InMemoryCommandBus struct {
	handlers map[reflect.Type]command.CommandHandlerInterface
}

func NewInMemoryCommandBus() *InMemoryCommandBus {
	return &InMemoryCommandBus{
		handlers: make(map[reflect.Type]command.CommandHandlerInterface),
	}
}

func (bus *InMemoryCommandBus) RegisterHandler(cmd command.CommandInterface, handler command.CommandHandlerInterface) {
	commandType := reflect.TypeOf(cmd)
	if _, exists := bus.handlers[commandType]; exists {
		panic("Command handler already registered for command type: " + commandType.String())
	}
	bus.handlers[commandType] = handler
}

func (bus *InMemoryCommandBus) Handle(command command.CommandInterface) error {
	commandType := reflect.TypeOf(command)
	_, exists := bus.handlers[commandType]
	if !exists {
		return exceptions.ErrQueryHandlerNotFound
	}
	return nil
}
