package bootstrap

import "GoAPISkyService/internal/flight_context/infrastructure/bootstrap/bus"

func RegisterCommands(commandBus *bus.InMemoryCommandBus, commandHandlers ...interface{}) {
	// TODO: comnplete RegisterCommands
	// for _, handler := range commandHandlers {
	// 	if cmdHandler, ok := handler.(bus.); ok {
	// 		commandBus.RegisterHandler(cmdHandler.CommandType(), cmdHandler)
	// 	} else {
	// 		panic("Invalid command handler type: " + handler.(string))
	// 	}
	// }) {
}