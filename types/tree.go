package types

type Tree struct {
	self Node

	decider  TreeDecider
	children []*Tree
}
