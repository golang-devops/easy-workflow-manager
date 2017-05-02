package coffee

import (
	"time"

	"github.com/golang-devops/easy-workflow-manager/attempt3"
)

type MixIngredientsFork struct {
	eventHandler *EventHandler
	legs         []attempt3.Task
	nextNode     attempt3.Node
}

func (m *MixIngredientsFork) Name() string {
	return "Mix Ingredients Fork"
}

func (m *MixIngredientsFork) Accept(visitor attempt3.NodeVisitor) {
	visitor.VisitParallelFork(m)
}

func (m *MixIngredientsFork) Legs() []attempt3.Task {
	return m.legs
}

func (m *MixIngredientsFork) Next() attempt3.Node {
	return m.nextNode
}

type AddMilk struct {
	eventHandler *EventHandler
}

func (a *AddMilk) Name() string {
	return "Add Milk"
}

func (a *AddMilk) Accept(visitor attempt3.NodeVisitor) {
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

func (s *StirMug) Accept(visitor attempt3.NodeVisitor) {
	visitor.VisitTask(s)
}

func (s *StirMug) Execute() error {
	s.eventHandler.Info("Start stirring Mug")
	time.Sleep(1 * time.Second)
	s.eventHandler.Info("Finish stirring Mug")
	return nil
}
