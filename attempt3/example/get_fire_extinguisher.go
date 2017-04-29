package example

import (
	"github.com/golang-devops/easy-workflow-manager/attempt3"
)

type GetFireExtinguisher struct {
	eventHandler *EventHandler
	nextNode     attempt3.Node
}

func (g *GetFireExtinguisher) Name() string {
	return "GetFireExtinguisher"
}

func (g *GetFireExtinguisher) Accept(visitor attempt3.NodeVisitor) {
	visitor.VisitActivity(g)
}

func (g *GetFireExtinguisher) Execute() error {
	g.eventHandler.Info("Extinguishing fire!")
	return nil
}

func (g *GetFireExtinguisher) Next() attempt3.Node {
	return g.nextNode
}
