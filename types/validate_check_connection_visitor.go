package types

import (
	"errors"
	"fmt"
)

type validateCheckConnectionVisitor struct {
	workflow       *Workflow
	processedNodes []Node
	nodeError      *NodeError
}

func (v *validateCheckConnectionVisitor) Validate() error {
	v.nodeError = nil

	v.workflow.initialNode.Accept(v)
	if v.nodeError != nil {
		return fmt.Errorf("Node '%s', %s", v.nodeError.Node.Name(), v.nodeError.Error.Error())
	}

	return nil
}
func (v *validateCheckConnectionVisitor) isAlreadyProcessed(node Node) bool {
	for _, n := range v.processedNodes {
		if n == node {
			return true
		}
	}
	return false
}
func (v *validateCheckConnectionVisitor) VisitTask(node Task) {
	if v.isAlreadyProcessed(node) {
		return
	}
	v.processedNodes = append(v.processedNodes, node)
}
func (v *validateCheckConnectionVisitor) VisitActivity(node Activity) {
	if v.isAlreadyProcessed(node) {
		return
	}
	v.processedNodes = append(v.processedNodes, node)

	if node.Next() != nil {
		node.Next().Accept(v)
		if v.nodeError != nil {
			return
		}
	}
}
func (v *validateCheckConnectionVisitor) VisitCondition(node Condition) {
	if v.isAlreadyProcessed(node) {
		return
	}
	v.processedNodes = append(v.processedNodes, node)

	if node.ConditionAnswerProvider() == nil {
		v.nodeError = &NodeError{
			Node:  node,
			Error: errors.New("ConditionAnswerProvider is not set"),
		}
		return
	}

	if node.WhenTrue() == nil {
		v.nodeError = &NodeError{
			Node:  node,
			Error: errors.New("WhenTrue is not set"),
		}
		return
	}

	if node.WhenFalse() == nil {
		v.nodeError = &NodeError{
			Node:  node,
			Error: errors.New("WhenFalse is not set"),
		}
		return
	}

	for _, child := range []Node{node.WhenTrue(), node.WhenFalse()} {
		child.Accept(v)
		if v.nodeError != nil {
			return
		}
	}
}
func (v *validateCheckConnectionVisitor) VisitSwitch(node Switch) {
	if v.isAlreadyProcessed(node) {
		return
	}
	v.processedNodes = append(v.processedNodes, node)

	if len(node.Cases()) == 0 {
		v.nodeError = &NodeError{
			Node:  node,
			Error: errors.New("No Cases are specified"),
		}
		return
	}

	uniqueNodes := map[Node]interface{}{}
	for _, switchCase := range node.Cases() {
		if _, alreadyFound := uniqueNodes[switchCase.Node]; alreadyFound {
			v.nodeError = &NodeError{
				Node:  node,
				Error: fmt.Errorf("Duplicate case '%s'", switchCase.Node.Name()),
			}
			return
		}
		uniqueNodes[switchCase.Node] = nil
	}

	for _, child := range node.Cases() {
		child.Node.Accept(v)
		if v.nodeError != nil {
			return
		}
	}
}
func (v *validateCheckConnectionVisitor) VisitParallelFork(node ParallelFork) {
	if v.isAlreadyProcessed(node) {
		return
	}
	v.processedNodes = append(v.processedNodes, node)

	if node.Next() != nil {
		nextNode := node.Next()
		nextNode.Accept(v)
		if v.nodeError != nil {
			return
		}
	}

	if len(node.Legs()) == 0 {
		v.nodeError = &NodeError{
			Node:  node,
			Error: errors.New("No Legs are specified"),
		}
		return
	}

	for _, child := range node.Legs() {
		child.Accept(v)
		if v.nodeError != nil {
			return
		}
	}
}
