package main

import (
	"fmt"

	"github.com/golang-devops/easy-workflow-manager/types"
)

type sharedData struct {
	data map[string]interface{}
}

func (s *sharedData) Set(name string, value interface{}) error {
	s.data[name] = value
}
func (s *sharedData) Get(name string) (interface{}, error) {
	val, ok := s.data[name]
	if !ok {
		return nil, fmt.Errorf("Value with name '%s' not found", name)
	}
	return val, nil
}

func main() {
	sd := &sharedData{}
	workflowTree := &types.Tree{}

	types.NewWorkflowExecutor(workflowTree, sd)
}
