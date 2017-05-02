package types

import (
	"errors"
	"fmt"

	"github.com/golang-devops/easy-workflow-manager/flowdiagrams"
)

type Workflow struct {
	eventHandler EventHandler
	initialNode  Node

	executedNodes []Node
}

func (w *Workflow) Validate() error {
	if (w.eventHandler) == nil {
		return errors.New("EventHandler is required")
	}

	if w.initialNode == nil {
		return errors.New("Workflow must have an Initial Node")
	}

	type Validator interface {
		Validate() error
	}

	validators := []Validator{
		&validateCheckConnectionVisitor{workflow: w},
	}

	for _, validator := range validators {
		if err := validator.Validate(); err != nil {
			return fmt.Errorf("Validation failed for validator %T, error: %s", validator, err.Error())
		}
	}

	return nil
}

func (w *Workflow) didNodeExecute(node Node) bool {
	for _, n := range w.executedNodes {
		if n == node {
			return true
		}
	}
	return false
}

func (w *Workflow) getDrawer(drawerForExecutedPath bool) *flowdiagrams.Drawer {
	drawer := flowdiagrams.NewDrawer("#333333")

	nodeIDs := make(map[Node]string)

	flattenedNodesVisitor := &getFlattenedNodesVisitor{}
	w.initialNode.Accept(flattenedNodesVisitor)

	neutralCategoryProvider := flowdiagrams.NewSimpleCategoryProvider("Neutral", "#FFDDDDDD", "#FFDDDDDD")
	executedCategoryProvider := flowdiagrams.NewSimpleCategoryProvider("Executed", "#FFDDDDDD", "#FFDDDDDD")
	notExecutedCategoryProvider := flowdiagrams.NewSimpleCategoryProvider("NotExecuted", "Transparent", "#FF888888")

	for _, node := range flattenedNodesVisitor.flattenedNodes {
		var categoryProvider flowdiagrams.CategoryProvider
		if !drawerForExecutedPath {
			categoryProvider = neutralCategoryProvider
		} else if w.didNodeExecute(node) {
			categoryProvider = executedCategoryProvider
		} else {
			categoryProvider = notExecutedCategoryProvider
		}

		nodeID := drawer.AddNode(node.Name(), categoryProvider)
		nodeIDs[node] = nodeID
	}

	for _, node := range flattenedNodesVisitor.flattenedNodes {
		nodeLinksVisitor := &nodeLinksToVisitor{workflow: w}
		node.Accept(nodeLinksVisitor)

		for _, link := range nodeLinksVisitor.linksTo {
			var categoryProvider flowdiagrams.CategoryProvider
			if !drawerForExecutedPath {
				categoryProvider = neutralCategoryProvider
			} else if w.didNodeExecute(link.FromNode) && w.didNodeExecute(link.ToNode) {
				categoryProvider = executedCategoryProvider
			} else {
				categoryProvider = notExecutedCategoryProvider
			}

			drawer.AddLink(nodeIDs[link.FromNode], nodeIDs[link.ToNode], link.LinkName, categoryProvider)
		}
	}

	return drawer
}

func (w *Workflow) DefaultDrawer() *flowdiagrams.Drawer {
	return w.getDrawer(false)
}

func (w *Workflow) ExecutedPathDrawer() *flowdiagrams.Drawer {
	return w.getDrawer(true)
}

func (w *Workflow) Execute() error {
	w.executedNodes = nil

	if err := w.Validate(); err != nil {
		return err
	}

	currentNode := w.initialNode
	for {
		w.eventHandler.Info(fmt.Sprintf("Executing: %s", currentNode.Name()))

		executeAndNext := &executeAndGetNextNodeVisitor{}
		currentNode.Accept(executeAndNext)
		if executeAndNext.executeErr != nil {
			w.eventHandler.Error(fmt.Sprintf("Failed to Execute: %s. Error: %s", currentNode.Name(), executeAndNext.executeErr.Error()))
			return executeAndNext.executeErr
		}
		w.eventHandler.Info(fmt.Sprintf("Successfully Executed: %s", currentNode.Name()))

		w.executedNodes = append(w.executedNodes, currentNode)
		w.executedNodes = append(w.executedNodes, executeAndNext.subNodesExecuted...)

		if executeAndNext.nextNode != nil {
			currentNode = executeAndNext.nextNode
		} else {
			w.eventHandler.Info(fmt.Sprintf("Node '%s' is the last node", currentNode.Name()))
			break
		}

	}

	return nil
}
