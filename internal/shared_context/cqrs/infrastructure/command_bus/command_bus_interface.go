package command_bus

import "GoAPISkyService/internal/shared_context/cqrs/application/command"

type CommandBusInterface interface {
	Handle(command command.CommandInterface) error
}
