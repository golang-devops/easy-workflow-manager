package attempt3

type ParallelFork interface {
	Node

	Legs() []Task
	Next() Node
}
