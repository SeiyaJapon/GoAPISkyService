package query_bus

import "GoAPISkyService/internal/shared_context/cqrs/application/query"

type QueryBusInterface interface {
	Ask(query query.QueryInterface) (result any, err error)
}
