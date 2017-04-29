package attempt3

import "fmt"

type validateReachabilityVisitor struct {
	workflow  *Workflow
	nodeError error
}

func (v *validateReachabilityVisitor) Validate() error {
	v.nodeError = nil
	for _, node := range v.workflow.nodes {
		node.Accept(v)
		if v.nodeError != nil {
			return fmt.Errorf("Node '%s', %s", node.Name(), v.nodeError.Error())
		}
	}
	return nil
}

func (v *validateReachabilityVisitor) checkNodeIsReachable(node Node) error {
	if v.workflow.InitialNode() == node {
		return nil
	}

	for _, to := range v.workflow.connectionsFromTo {
		if to == node {
			return nil
		}
	}

	for _, node2 := range v.workflow.nodes {
		if node2 == node {
			continue
		}
		visitor := &nodeContainsAnotherNodeVisitor{nodeToContain: node}
		node2.Accept(visitor)
		if visitor.containsIt {
			return nil
		}
	}

	return fmt.Errorf("Node '%s' is unreachable", node.Name())
}
func (v *validateReachabilityVisitor) VisitActivity(node Activity) {
	if err := v.checkNodeIsReachable(node); err != nil {
		v.nodeError = err
		return
	}
}
func (v *validateReachabilityVisitor) VisitCondition(node Condition) {
	if err := v.checkNodeIsReachable(node); err != nil {
		v.nodeError = err
		return
	}
}
func (v *validateReachabilityVisitor) VisitSwitch(node Switch) {
	if err := v.checkNodeIsReachable(node); err != nil {
		v.nodeError = err
		return
	}
}
func (v *validateReachabilityVisitor) VisitParallelFork(node ParallelFork) {
	if err := v.checkNodeIsReachable(node); err != nil {
		v.nodeError = err
		return
	}
}
