package attempt3

type Activity interface {
	Node

	Execute() error
}
