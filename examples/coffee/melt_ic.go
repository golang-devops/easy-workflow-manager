package coffee

import (
	"github.com/golang-devops/easy-workflow-manager/types"
	"github.com/golang-devops/easy-workflow-manager/logging"
)

type MeltIce struct {
	eventHandler logging.Logger
	nextNode     types.Node
}

func (m *MeltIce) Name() string {
	return "MeltIce"
}

func (m *MeltIce) Accept(visitor types.NodeVisitor) {
	visitor.VisitActivity(m)
}

func (m *MeltIce) Execute() error {
	m.eventHandler.Info("Extinguishing fire!")
	return nil
}

func (m *MeltIce) Next() types.Node {
	return m.nextNode
}
