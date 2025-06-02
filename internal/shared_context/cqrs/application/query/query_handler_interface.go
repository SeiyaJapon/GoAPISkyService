package query

type QueryHandlerInterface interface {
	Handle(query QueryInterface) (result QueryResultInterface, err error)
}
