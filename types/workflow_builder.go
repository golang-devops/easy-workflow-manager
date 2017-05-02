package types

type WorkflowBuilder struct {
	workflow *Workflow
}

func NewWorkflowBuilder(initialNode Node) *WorkflowBuilder {
	return &WorkflowBuilder{
		workflow: &Workflow{
			initialNode: initialNode,
		},
	}
}

func (w *WorkflowBuilder) SetEventHandler(eventHandler EventHandler) *WorkflowBuilder {
	w.workflow.eventHandler = eventHandler
	return w
}

func (w *WorkflowBuilder) Build() (*Workflow, error) {
	if err := w.workflow.Validate(); err != nil {
		return nil, err
	}
	return w.workflow, nil
}
