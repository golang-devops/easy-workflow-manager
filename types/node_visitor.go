package types

type NodeVisitor interface {
	VisitTask(node Task)
	VisitActivity(node Activity)
	VisitCondition(node Condition)
	VisitSwitch(node Switch)
	VisitParallelFork(node ParallelFork)
}
