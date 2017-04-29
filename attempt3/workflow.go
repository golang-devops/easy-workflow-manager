package attempt3

import (
	"errors"
	"fmt"

	"github.com/golang-devops/easy-workflow-manager/flowdiagrams"
)

type Workflow struct {
	eventHandler      EventHandler
	nodes             []Node
	connectionsFromTo map[Node]Node
}

func (w *Workflow) InitialNode() Node {
	return w.nodes[0]
}

func (w *Workflow) Validate() error {
	if (w.eventHandler) == nil {
		return errors.New("EventHandler is required")
	}

	if len(w.nodes) == 0 {
		return errors.New("Workflow must have at least one Node")
	}

	type Validator interface {
		Validate() error
	}

	validators := []Validator{
		&validateCheckConnectionVisitor{workflow: w},
		&validateReachabilityVisitor{workflow: w},
	}

	for _, validator := range validators {
		if err := validator.Validate(); err != nil {
			return fmt.Errorf("Validation failed for validator %T, error: %s", validator, err.Error())
		}
	}

	return nil
}

func (w *Workflow) Drawer() *flowdiagrams.Drawer {
	drawer := flowdiagrams.NewDrawer()

	for _, node := range w.nodes {
		drawer.AddNode(node.Name())
	}

	return drawer
}

func (w *Workflow) Execute() error {
	if err := w.Validate(); err != nil {
		return err
	}

	currentNode := w.InitialNode()
	for {
		err := executeNode(currentNode)
		if err != nil {
			return err
		}

		nextNodeVisitor := &getNextNodeVisitor{}
		currentNode.Accept(nextNodeVisitor)
		if nextNodeVisitor.nextNode != nil {
			currentNode = nextNodeVisitor.nextNode
		} else {
			tmpCurrentNodeCopy := currentNode
			tmpNextCurrentNode, ok := w.connectionsFromTo[tmpCurrentNodeCopy]
			if !ok {
				w.eventHandler.Info(fmt.Sprintf("Node '%s' is the last node", tmpCurrentNodeCopy.Name()))
				break
			}
			currentNode = tmpNextCurrentNode
		}

	}

	return nil
}
