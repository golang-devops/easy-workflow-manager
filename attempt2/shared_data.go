package attempt2

type SharedData interface {
	Set(name string, value interface{}) error
	Get(name string) (interface{}, error)
}
