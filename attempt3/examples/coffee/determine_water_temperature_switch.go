package coffee

import (
	"github.com/golang-devops/easy-workflow-manager/attempt3"
)

type DetermineWaterTemperatureSwitch struct {
	eventHandler *EventHandler
	cases        []*attempt3.SwitchCase
}

func (d *DetermineWaterTemperatureSwitch) Name() string {
	return "Determine Temperature"
}

func (d *DetermineWaterTemperatureSwitch) Cases() []*attempt3.SwitchCase {
	return d.cases
}

func (d *DetermineWaterTemperatureSwitch) Accept(visitor attempt3.NodeVisitor) {
	visitor.VisitSwitch(d)
}
