package example

import (
	"fmt"
)

type EventHandler struct{}

func (e *EventHandler) Info(msg string) {
	fmt.Println(msg)
}
