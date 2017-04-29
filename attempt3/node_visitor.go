package attempt3

type NodeVisitor interface {
	VisitActivity(node Activity)
	VisitCondition(node Condition)
	VisitSwitch(node Switch)
	VisitParallelFork(node ParallelFork)
}
