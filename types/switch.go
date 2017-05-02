package types

type Switch interface {
	Node

	Cases() []*SwitchCase
}

type SwitchCase struct {
	Predicate func() bool
	Name      string
	Node      Node
}
