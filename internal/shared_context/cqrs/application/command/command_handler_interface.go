package command

type CommandHandlerInterface interface {
	Handler(command CommandInterface) error
}
