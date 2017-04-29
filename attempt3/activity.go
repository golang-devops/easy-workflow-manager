package attempt3

type Activity interface {
	Task

	Next() Node
}
