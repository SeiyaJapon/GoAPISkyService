package command

type CommandHandlerInterface interface {
	Handle(command CommandInterface) error
}
