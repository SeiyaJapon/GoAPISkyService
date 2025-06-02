package query

type QueryResultInterface interface {
	result() (slice any, err error)
}
