package types

type ParallelFork interface {
	Node

	Legs() []Task
	Next() Node
}
