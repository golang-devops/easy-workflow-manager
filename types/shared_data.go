package types

type SharedData interface {
	Set(name string, value interface{}) error
	Get(name string) (interface{}, error)
}
