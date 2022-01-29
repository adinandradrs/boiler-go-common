package base

type BaseService interface {
	Execute(input interface{}) (interface{}, error)
}
