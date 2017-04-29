package attempt2

type Tree struct {
	self Node

	decider  TreeDecider
	children []*Tree
}
