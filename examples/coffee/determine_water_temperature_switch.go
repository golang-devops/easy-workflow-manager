package coffee

import (
	"github.com/golang-devops/easy-workflow-manager/types"
	"github.com/golang-devops/easy-workflow-manager/logging"
)

type DetermineWaterTemperatureSwitch struct {
	eventHandler logging.Logger
	cases        []*types.SwitchCase
}

func (d *DetermineWaterTemperatureSwitch) Name() string {
	return "Determine Temperature"
}

func (d *DetermineWaterTemperatureSwitch) Cases() []*types.SwitchCase {
	return d.cases
}

func (d *DetermineWaterTemperatureSwitch) Accept(visitor types.NodeVisitor) {
	visitor.VisitSwitch(d)
}
