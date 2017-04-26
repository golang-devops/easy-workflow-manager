package elements

type Activity interface {
	Execute() error
}

type Task interface {
	Activity
}

type TaskWithRollback interface {
	Task
	Rollback() error
}

//TODO: Implement -> type SubProcessActivity /*AKA "compound activity"*/ struct{}

type Transaction struct {
	Tasks []*TaskWithRollback
}

type CallActivity interface {
	Task
}
