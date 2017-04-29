package main

import (
	"fmt"
	"log"

	"time"

	"github.com/golang-devops/easy-workflow-manager/attempt2"
	"github.com/golang-devops/easy-workflow-manager/attempt3"
)

type sharedData struct {
	data map[string]interface{}
}

func (s *sharedData) Set(name string, value interface{}) error {
	s.data[name] = value
	return nil
}
func (s *sharedData) Get(name string) (interface{}, error) {
	val, ok := s.data[name]
	if !ok {
		return nil, fmt.Errorf("Value with name '%s' not found", name)
	}
	return val, nil
}

func tmpAttempt2() {
	sd := &sharedData{}
	workflowTree := &attempt2.Tree{}

	sd = sd
	workflowTree = workflowTree
	// attempt2.NewWorkflowExecutor(workflowTree, sd)
}

type EventHandler struct{}

func (e *EventHandler) Info(msg string) {
	fmt.Println(msg)
}

type HumanBoilKettleActivity struct {
	eventHandler *EventHandler
}

func (h *HumanBoilKettleActivity) Name() string {
	return "Boil Kettle"
}
func (h *HumanBoilKettleActivity) Execute() error {
	h.eventHandler.Info("Water starting to boil")
	time.Sleep(2 * time.Second)
	h.eventHandler.Info("Water finished boiling")
	return nil //TODO:
}
func (h *HumanBoilKettleActivity) Accept(visitor attempt3.NodeVisitor) {
	visitor.VisitActivity(h)
}

type DetermineWaterTemperatureSwitch struct {
	eventHandler *EventHandler
}

func (d *DetermineWaterTemperatureSwitch) Name() string {
	return "Determine Temperature"
}
func (d *DetermineWaterTemperatureSwitch) Accept(visitor attempt3.NodeVisitor) {
	visitor.VisitSwitch(d)
}
func (d *DetermineWaterTemperatureSwitch) SwitchAnswerProvider() attempt3.SwitchAnswerProvider {
	return d
}
func (d *DetermineWaterTemperatureSwitch) AllCases() []attempt3.Node {
	return nil //TODO:
}
func (d *DetermineWaterTemperatureSwitch) SwitchAnswer() attempt3.Node {
	return nil //TODO:
}

func tmpAttempt3() {
	eventHandler := &EventHandler{}

	var (
		humanBoilKettle           attempt3.Activity = &HumanBoilKettleActivity{eventHandler: eventHandler}
		determineWaterTemperature attempt3.Switch   = &DetermineWaterTemperatureSwitch{eventHandler: eventHandler}
	)

	coffeeWorkflow, err := attempt3.NewWorkflowBuilder().
		SetEventHandler(eventHandler).
		AddNode(humanBoilKettle).
		AddNode(determineWaterTemperature).
		ConnectNodes(humanBoilKettle, determineWaterTemperature).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	/*drawer := coffeeWorkflow.Drawer()
	if err := drawer.SaveToXml("sample_workflow.xml"); err != nil {
		log.Fatal(err)
	}*/

	if err := coffeeWorkflow.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	tmpAttempt2()
	tmpAttempt3()
}
