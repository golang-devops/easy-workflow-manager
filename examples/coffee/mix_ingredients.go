package coffee

import (
	"time"

	"github.com/golang-devops/easy-workflow-manager/types"
)

type MixIngredientsFork struct {
	eventHandler *EventHandler
	legs         []types.Task
	nextNode     types.Node
}

func (m *MixIngredientsFork) Name() string {
	return "Mix Ingredients Fork"
}

func (m *MixIngredientsFork) Accept(visitor types.NodeVisitor) {
	visitor.VisitParallelFork(m)
}

func (m *MixIngredientsFork) Legs() []types.Task {
	return m.legs
}

func (m *MixIngredientsFork) Next() types.Node {
	return m.nextNode
}

type AddMilk struct {
	eventHandler *EventHandler
}

func (a *AddMilk) Name() string {
	return "Add Milk"
}

func (a *AddMilk) Accept(visitor types.NodeVisitor) {
	visitor.VisitTask(a)
}

func (a *AddMilk) Execute() error {
	a.eventHandler.Info("Start adding Milk")
	time.Sleep(1 * time.Second)
	a.eventHandler.Info("Finish adding Milk")
	return nil
}

type StirMug struct {
	eventHandler *EventHandler
}

func (s *StirMug) Name() string {
	return "Stir Mug"
}

func (s *StirMug) Accept(visitor types.NodeVisitor) {
	visitor.VisitTask(s)
}

func (s *StirMug) Execute() error {
	s.eventHandler.Info("Start stirring Mug")
	time.Sleep(1 * time.Second)
	s.eventHandler.Info("Finish stirring Mug")
	return nil
}
