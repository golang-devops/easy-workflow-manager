package types

type Activity interface {
	Task

	Next() Node
}
