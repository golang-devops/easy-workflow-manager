package types

import (
	"errors"
	"fmt"
	"strings"
)

type executeAndGetNextNodeVisitor struct {
	executeErr       error
	nextNode         Node
	subNodesExecuted []Node
}

func (e *executeAndGetNextNodeVisitor) VisitTask(node Task) {
	e.executeErr = errors.New("Tasks should not get called via the executeAndGetNextNodeVisitor")
}
func (e *executeAndGetNextNodeVisitor) VisitActivity(node Activity) {
	e.executeErr = executeTask(node)
	if e.executeErr == nil {
		e.nextNode = node.Next()
	}
}

func (e *executeAndGetNextNodeVisitor) VisitCondition(node Condition) {
	if node.ConditionAnswerProvider().ConditionAnswer() {
		e.nextNode = node.WhenTrue()
	} else {
		e.nextNode = node.WhenFalse()
	}
}

func (e *executeAndGetNextNodeVisitor) VisitSwitch(node Switch) {
	for _, switchCase := range node.Cases() {
		if switchCase.Predicate() {
			e.nextNode = switchCase.Node
			return
		}
	}

	caseNames := []string{}
	for _, switchCase := range node.Cases() {
		caseNames = append(caseNames, switchCase.Name)
	}

	e.executeErr = fmt.Errorf("All Switch Cases were false (%s)", strings.Join(caseNames, " & "))
}

func (e *executeAndGetNextNodeVisitor) VisitParallelFork(node ParallelFork) {
	legs := node.Legs()

	//code inspired by http://www.golangpatterns.info/concurrency/parallel-for-loop
	type result struct {
		Error error
	}
	sem := make(chan *result, len(legs)) // semaphore pattern
	for _, leg := range legs {
		go func(copiedLeg Task) {
			res := &result{}
			res.Error = executeTask(copiedLeg)
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
		e.executeErr = errors.New(strings.Join(errorStrings, " & "))
		return
	}

	for _, leg := range legs {
		e.subNodesExecuted = append(e.subNodesExecuted, leg)
	}
	e.nextNode = node.Next()
}

type nodeContainsAnotherNodeVisitor struct {
	nodeToContain Node
	containsIt    bool
}

func (n *nodeContainsAnotherNodeVisitor) VisitTask(node Task)         {}
func (n *nodeContainsAnotherNodeVisitor) VisitActivity(node Activity) {}
func (n *nodeContainsAnotherNodeVisitor) VisitCondition(node Condition) {
	n.containsIt = node.WhenTrue() == n.nodeToContain || node.WhenFalse() == n.nodeToContain
}
func (n *nodeContainsAnotherNodeVisitor) VisitSwitch(node Switch) {
	for _, switchCase := range node.Cases() {
		if switchCase.Node == n.nodeToContain {
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

func executeTask(task Task) error {
	return task.Execute()
}

func getFlattenedNodes(initialNode Node) []Node {
	getFlattenedNodes := &getFlattenedNodesVisitor{}
	initialNode.Accept(getFlattenedNodes)
	return getFlattenedNodes.flattenedNodes
}

type getFlattenedNodesVisitor struct {
	flattenedNodes []Node
}

func (g *getFlattenedNodesVisitor) isAlreadyProcessed(node Node) bool {
	for _, n := range g.flattenedNodes {
		if n == node {
			return true
		}
	}
	return false
}
func (g *getFlattenedNodesVisitor) VisitTask(node Task) {
	if g.isAlreadyProcessed(node) {
		return
	}
	g.flattenedNodes = append(g.flattenedNodes, node)
}
func (g *getFlattenedNodesVisitor) VisitActivity(node Activity) {
	if g.isAlreadyProcessed(node) {
		return
	}
	g.flattenedNodes = append(g.flattenedNodes, node)

	if node.Next() != nil {
		node.Next().Accept(g)
	}
}
func (g *getFlattenedNodesVisitor) VisitCondition(node Condition) {
	if g.isAlreadyProcessed(node) {
		return
	}
	g.flattenedNodes = append(g.flattenedNodes, node)

	for _, child := range []Node{node.WhenTrue(), node.WhenFalse()} {
		child.Accept(g)
	}
}
func (g *getFlattenedNodesVisitor) VisitSwitch(node Switch) {
	if g.isAlreadyProcessed(node) {
		return
	}
	g.flattenedNodes = append(g.flattenedNodes, node)

	for _, child := range node.Cases() {
		child.Node.Accept(g)
	}
}
func (g *getFlattenedNodesVisitor) VisitParallelFork(node ParallelFork) {
	if g.isAlreadyProcessed(node) {
		return
	}
	g.flattenedNodes = append(g.flattenedNodes, node)

	for _, child := range node.Legs() {
		child.Accept(g)
	}

	if node.Next() != nil {
		node.Next().Accept(g)
	}
}

type nodeLink struct {
	LinkName string
	FromNode Node
	ToNode   Node
}

type nodeLinksToVisitor struct {
	workflow *Workflow
	linksTo  []*nodeLink
}

func (n *nodeLinksToVisitor) VisitTask(node Task) {
}
func (n *nodeLinksToVisitor) VisitActivity(node Activity) {
	if node.Next() != nil {
		n.linksTo = append(n.linksTo, &nodeLink{
			LinkName: "Direct",
			FromNode: node,
			ToNode:   node.Next(),
		})
	}
}
func (n *nodeLinksToVisitor) VisitCondition(node Condition) {
	n.linksTo = append(n.linksTo, &nodeLink{
		LinkName: "WhenTrue",
		FromNode: node,
		ToNode:   node.WhenTrue(),
	})
	n.linksTo = append(n.linksTo, &nodeLink{
		LinkName: "WhenFalse",
		FromNode: node,
		ToNode:   node.WhenFalse(),
	})
}
func (n *nodeLinksToVisitor) VisitSwitch(node Switch) {
	for _, switchCase := range node.Cases() {
		n.linksTo = append(n.linksTo, &nodeLink{
			LinkName: switchCase.Name,
			FromNode: node,
			ToNode:   switchCase.Node,
		})
	}
}
func (n *nodeLinksToVisitor) VisitParallelFork(node ParallelFork) {
	for _, leg := range node.Legs() {
		n.linksTo = append(n.linksTo, &nodeLink{
			LinkName: "Parallel",
			FromNode: node,
			ToNode:   leg,
		})
	}

	if node.Next() != nil {
		for _, leg := range node.Legs() {
			n.linksTo = append(n.linksTo, &nodeLink{
				LinkName: "Direct",
				FromNode: leg,
				ToNode:   node.Next(),
			})
		}
	}
}
