package types

type Task interface {
	Node

	Execute() error
}
