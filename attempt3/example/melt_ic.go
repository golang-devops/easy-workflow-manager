package example

import (
	"github.com/golang-devops/easy-workflow-manager/attempt3"
)

type MeltIce struct {
	eventHandler *EventHandler
	nextNode     attempt3.Node
}

func (m *MeltIce) Name() string {
	return "MeltIce"
}

func (m *MeltIce) Accept(visitor attempt3.NodeVisitor) {
	visitor.VisitActivity(m)
}

func (m *MeltIce) Execute() error {
	m.eventHandler.Info("Extinguishing fire!")
	return nil
}

func (m *MeltIce) Next() attempt3.Node {
	return m.nextNode
}
