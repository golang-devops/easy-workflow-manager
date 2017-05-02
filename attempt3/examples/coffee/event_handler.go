package coffee

import (
	"fmt"
)

type EventHandler struct{}

func (e *EventHandler) Info(msg string) {
	fmt.Println(msg)
}

func (e *EventHandler) Error(msg string) {
	fmt.Println("ERROR:" + msg)
}
