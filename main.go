package main

import (
	"fmt"
	"log"

	"github.com/golang-devops/easy-workflow-manager/attempt3/example"
)

type sharedData struct {
	data map[string]interface{}
}

func (s *sharedData) Set(name string, value interface{}) error {
	s.data[name] = value
	return nil
}
func (s *sharedData) Get(name string) (interface{}, error) {
	val, ok := s.data[name]
	if !ok {
		return nil, fmt.Errorf("Value with name '%s' not found", name)
	}
	return val, nil
}

func tmpAttempt3() {
	if err := example.ExecuteWorkflowExample(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	tmpAttempt3()
}
