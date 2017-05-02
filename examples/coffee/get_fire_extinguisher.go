package coffee

import (
	"github.com/golang-devops/easy-workflow-manager/types"
	"github.com/golang-devops/easy-workflow-manager/logging"
)

type GetFireExtinguisher struct {
	eventHandler logging.Logger
	nextNode     types.Node
}

func (g *GetFireExtinguisher) Name() string {
	return "GetFireExtinguisher"
}

func (g *GetFireExtinguisher) Accept(visitor types.NodeVisitor) {
	visitor.VisitActivity(g)
}

func (g *GetFireExtinguisher) Execute() error {
	g.eventHandler.Info("Extinguishing fire!")
	return nil
}

func (g *GetFireExtinguisher) Next() types.Node {
	return g.nextNode
}
