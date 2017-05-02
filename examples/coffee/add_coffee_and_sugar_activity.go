package coffee

import (
	"github.com/golang-devops/easy-workflow-manager/types"
)

type AddCoffeeAndSugar struct {
	eventHandler *EventHandler
	nextNode     types.Node
}

func (a *AddCoffeeAndSugar) Name() string {
	return "Add Coffee and Sugar"
}

func (a *AddCoffeeAndSugar) Accept(visitor types.NodeVisitor) {
	visitor.VisitActivity(a)
}

func (a *AddCoffeeAndSugar) Execute() error {
	a.eventHandler.Info("Coffee and Sugar being added")
	return nil
}

func (a *AddCoffeeAndSugar) Next() types.Node {
	return a.nextNode
}
