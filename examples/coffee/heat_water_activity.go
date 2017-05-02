package coffee

import (
	"fmt"
	"time"

	"github.com/golang-devops/easy-workflow-manager/types"
	"github.com/golang-devops/easy-workflow-manager/logging"
)

type HeatWaterActivity struct {
	eventHandler     logging.Logger
	nextNode         types.Node
	waterTemperature int
}

func (h *HeatWaterActivity) Name() string {
	return "Heat Water"
}

func (h *HeatWaterActivity) Execute() error {
	h.eventHandler.Info(fmt.Sprintf("Heating up (from %d degrees)", h.waterTemperature))
	time.Sleep(800 * time.Millisecond)
	h.waterTemperature += 50
	h.eventHandler.Info(fmt.Sprintf("Turning off the heat (now at %d degrees)", h.waterTemperature))
	return nil
}

func (h *HeatWaterActivity) Accept(visitor types.NodeVisitor) {
	visitor.VisitActivity(h)
}

func (h *HeatWaterActivity) Next() types.Node {
	return h.nextNode
}

func (h *HeatWaterActivity) WaterTemperature() int {
	return h.waterTemperature
}
