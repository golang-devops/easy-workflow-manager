package attempt3

import (
	"errors"
	"strings"
)

type getNextNodeVisitor struct {
	nextNode Node
}

func (g *getNextNodeVisitor) VisitActivity(node Activity) {}

func (g *getNextNodeVisitor) VisitCondition(node Condition) {
	if node.ConditionAnswerProvider().ConditionAnswer() {
		g.nextNode = node.WhenTrue()
	} else {
		g.nextNode = node.WhenFalse()
	}
}

func (g *getNextNodeVisitor) VisitSwitch(node Switch) {
	g.nextNode = node.SwitchAnswerProvider().SwitchAnswer()
}

func (g *getNextNodeVisitor) VisitParallelFork(node ParallelFork) {}

type nodeContainsAnotherNodeVisitor struct {
	nodeToContain Node
	containsIt    bool
}

func (n *nodeContainsAnotherNodeVisitor) VisitActivity(node Activity) {}
func (n *nodeContainsAnotherNodeVisitor) VisitCondition(node Condition) {
	n.containsIt = node.WhenTrue() == n.nodeToContain || node.WhenFalse() == n.nodeToContain
}
func (n *nodeContainsAnotherNodeVisitor) VisitSwitch(node Switch) {
	for _, caseNode := range node.SwitchAnswerProvider().AllCases() {
		if caseNode == n.nodeToContain {
			n.containsIt = true
			return
		}
	}
}
func (n *nodeContainsAnotherNodeVisitor) VisitParallelFork(node ParallelFork) {
	for _, leg := range node.Legs() {
		if leg == n.nodeToContain {
			n.containsIt = true
			return
		}
	}
}

type executeNodeVisitor struct {
	err error
}

func (e *executeNodeVisitor) VisitActivity(node Activity) {
	e.err = node.Execute()
}

func (e *executeNodeVisitor) VisitCondition(node Condition) {}

func (e *executeNodeVisitor) VisitSwitch(node Switch) {}

func (e *executeNodeVisitor) VisitParallelFork(node ParallelFork) {
	legs := node.Legs()

	//code inspired by http://www.golangpatterns.info/concurrency/parallel-for-loop
	type result struct {
		Error error
	}
	sem := make(chan *result, len(legs)) // semaphore pattern
	for _, leg := range legs {
		go func(copiedLeg Activity) {
			res := &result{}
			res.Error = executeNode(copiedLeg)
			sem <- res
		}(leg)
	}

	results := []*result{}
	// wait for goroutines to finish
	for i := 0; i < len(legs); i++ {
		results = append(results, <-sem)
	}

	errorStrings := []string{}
	for _, res := range results {
		if res.Error != nil {
			errorStrings = append(errorStrings, res.Error.Error())
		}
	}

	if len(errorStrings) > 0 {
		e.err = errors.New(strings.Join(errorStrings, " & "))
	}
}

func executeNode(node Node) error {
	visitor := &executeNodeVisitor{}
	node.Accept(visitor)
	return visitor.err
}
