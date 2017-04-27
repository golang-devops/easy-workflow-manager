package types

type TreeDecider interface {
	Decide(sharedData SharedData, children []*Tree) *Tree
}
