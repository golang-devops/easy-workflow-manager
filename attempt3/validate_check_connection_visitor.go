package attempt3

import (
	"errors"
	"fmt"
)

type validateCheckConnectionVisitor struct {
	workflow  *Workflow
	nodeError error
}

func (v *validateCheckConnectionVisitor) Validate() error {
	v.nodeError = nil
	for _, node := range v.workflow.nodes {
		node.Accept(v)
		if v.nodeError != nil {
			return fmt.Errorf("Node '%s', %s", node.Name(), v.nodeError.Error())
		}
	}
	return nil
}
func (v *validateCheckConnectionVisitor) VisitActivity(node Activity) {
}
func (v *validateCheckConnectionVisitor) VisitCondition(node Condition) {
	if node.ConditionAnswerProvider() == nil {
		v.nodeError = errors.New("ConditionAnswerProvider is not set")
		return
	}

	if node.WhenTrue() == nil {
		v.nodeError = errors.New("WhenTrue is not set")
		return
	}

	if node.WhenFalse() == nil {
		v.nodeError = errors.New("WhenFalse is not set")
		return
	}
}
func (v *validateCheckConnectionVisitor) VisitSwitch(node Switch) {
	if node.SwitchAnswerProvider() == nil {
		v.nodeError = errors.New("SwitchAnswerProvider is not set")
		return
	}

	uniqueMap := map[Node]interface{}{}
	for _, c := range node.SwitchAnswerProvider().AllCases() {
		if _, alreadyFound := uniqueMap[c]; alreadyFound {
			v.nodeError = fmt.Errorf("Duplicate case '%s'", c.Name())
			return
		}
		uniqueMap[c] = nil
	}
}
func (v *validateCheckConnectionVisitor) VisitParallelFork(node ParallelFork) {}
