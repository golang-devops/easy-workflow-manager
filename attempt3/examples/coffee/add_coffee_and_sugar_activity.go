package coffee

import (
	"github.com/golang-devops/easy-workflow-manager/attempt3"
)

type AddCoffeeAndSugar struct {
	eventHandler *EventHandler
	nextNode     attempt3.Node
}

func (a *AddCoffeeAndSugar) Name() string {
	return "Add Coffee and Sugar"
}

func (a *AddCoffeeAndSugar) Accept(visitor attempt3.NodeVisitor) {
	visitor.VisitActivity(a)
}

func (a *AddCoffeeAndSugar) Execute() error {
	a.eventHandler.Info("Coffee and Sugar being added")
	return nil
}

func (a *AddCoffeeAndSugar) Next() attempt3.Node {
	return a.nextNode
}
