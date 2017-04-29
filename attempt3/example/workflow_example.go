package example

import (
	"github.com/golang-devops/easy-workflow-manager/attempt3"
)

func ExecuteWorkflowExample() error {
	var (
		eventHandler = &EventHandler{}

		heatWater = &HeatWaterActivity{
			eventHandler: eventHandler,
		}

		addCoffeeAndSugar = &AddCoffeeAndSugar{
			eventHandler: eventHandler,
		}

		mixIngredientsFork = &MixIngredientsFork{
			eventHandler: eventHandler,
			legs: []attempt3.Task{
				&AddMilk{eventHandler: eventHandler},
				&StirMug{eventHandler: eventHandler},
			},
		}

		drinkCoffee = &DrinkCoffee{eventHandler: eventHandler}

		meltIce             = &MeltIce{eventHandler: eventHandler}
		getFireExtinguisher = &GetFireExtinguisher{eventHandler: eventHandler}

		determineWaterTemperature attempt3.Switch = &DetermineWaterTemperatureSwitch{
			eventHandler: eventHandler,
			cases: []*attempt3.SwitchCase{
				&attempt3.SwitchCase{
					Name:      "Temperature <= 0 degrees",
					Predicate: func() bool { return heatWater.WaterTemperature() <= 0 },
					Node:      meltIce,
				},
				&attempt3.SwitchCase{
					Name:      "Temperature > 0 & Temperature < 100 degrees",
					Predicate: func() bool { return heatWater.WaterTemperature() > 0 && heatWater.WaterTemperature() < 100 },
					Node:      heatWater,
				},
				&attempt3.SwitchCase{
					Name:      "Temperature > 100 & Temperature < 160 degrees",
					Predicate: func() bool { return heatWater.WaterTemperature() < 160 },
					Node:      addCoffeeAndSugar,
				},
				&attempt3.SwitchCase{
					Name:      "Temperature >= 160 degrees",
					Predicate: func() bool { return heatWater.WaterTemperature() >= 160 },
					Node:      getFireExtinguisher,
				},
			},
		}
	)

	heatWater.nextNode = determineWaterTemperature
	meltIce.nextNode = determineWaterTemperature
	getFireExtinguisher.nextNode = determineWaterTemperature
	addCoffeeAndSugar.nextNode = mixIngredientsFork
	mixIngredientsFork.nextNode = drinkCoffee

	coffeeWorkflow, err := attempt3.NewWorkflowBuilder(heatWater).
		SetEventHandler(eventHandler).
		Build()
	if err != nil {
		return err
	}

	nonExecutedDrawer := coffeeWorkflow.DefaultDrawer()
	if err := nonExecutedDrawer.SaveToDGML("sample_workflow.dgml"); err != nil {
		return err
	}

	if err := coffeeWorkflow.Execute(); err != nil {
		return err
	}

	executedDrawer := coffeeWorkflow.ExecutedPathDrawer()
	if err := executedDrawer.SaveToDGML("sample_workflow_executed.dgml"); err != nil {
		return err
	}

	return nil
}
