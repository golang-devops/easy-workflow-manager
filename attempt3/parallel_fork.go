package attempt3

type ParallelFork interface {
	Node

	AddLeg(activity Activity) ParallelFork

	Legs() []Activity
}
