package attempt3

type Task interface {
	Node

	Execute() error
}
