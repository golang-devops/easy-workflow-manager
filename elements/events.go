package elements

type Event struct{}
type StartEvent struct{ Event }
type IntermediateEvent struct{ Event }
type EndEvent struct{ Event }