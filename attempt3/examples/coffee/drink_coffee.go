package coffee

import (
	"time"

	"github.com/golang-devops/easy-workflow-manager/attempt3"
)

type DrinkCoffee struct {
	eventHandler *EventHandler
}

func (d *DrinkCoffee) Name() string {
	return "DrinkCoffee"
}

func (d *DrinkCoffee) Accept(visitor attempt3.NodeVisitor) {
	visitor.VisitActivity(d)
}

func (d *DrinkCoffee) Execute() error {
	d.eventHandler.Info("Drinking coffee...")
	time.Sleep(500 * time.Millisecond)
	d.eventHandler.Info("Drinking coffee...")
	time.Sleep(500 * time.Millisecond)
	d.eventHandler.Info("Drinking coffee...")
	time.Sleep(500 * time.Millisecond)
	d.eventHandler.Info("What a lovely cup!")
	return nil
}

func (d *DrinkCoffee) Next() attempt3.Node {
	return nil
}
