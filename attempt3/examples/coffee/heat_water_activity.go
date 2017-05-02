package coffee

import (
	"fmt"
	"time"

	"github.com/golang-devops/easy-workflow-manager/attempt3"
)

type HeatWaterActivity struct {
	eventHandler     *EventHandler
	nextNode         attempt3.Node
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

func (h *HeatWaterActivity) Accept(visitor attempt3.NodeVisitor) {
	visitor.VisitActivity(h)
}

func (h *HeatWaterActivity) Next() attempt3.Node {
	return h.nextNode
}

func (h *HeatWaterActivity) WaterTemperature() int {
	return h.waterTemperature
}
