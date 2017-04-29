package attempt2

type TreeDecider interface {
	Decide(sharedData SharedData, children []*Tree) *Tree
}
