package coffee

import (
	"time"

	"github.com/golang-devops/easy-workflow-manager/types"
	"github.com/golang-devops/easy-workflow-manager/logging"
)

type DrinkCoffee struct {
	eventHandler logging.Logger
}

func (d *DrinkCoffee) Name() string {
	return "DrinkCoffee"
}

func (d *DrinkCoffee) Accept(visitor types.NodeVisitor) {
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

func (d *DrinkCoffee) Next() types.Node {
	return nil
}
