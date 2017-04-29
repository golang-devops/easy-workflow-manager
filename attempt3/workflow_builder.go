package attempt3

type WorkflowBuilder struct {
	workflow *Workflow
}

func NewWorkflowBuilder() *WorkflowBuilder {
	return &WorkflowBuilder{
		workflow: &Workflow{
			connectionsFromTo: make(map[Node]Node),
		},
	}
}

func (w *WorkflowBuilder) SetEventHandler(eventHandler EventHandler) *WorkflowBuilder {
	w.workflow.eventHandler = eventHandler
	return w
}

func (w *WorkflowBuilder) AddNode(node Node) *WorkflowBuilder {
	w.workflow.nodes = append(w.workflow.nodes, node)
	return w
}

func (w *WorkflowBuilder) ConnectNodes(from, to Node) *WorkflowBuilder {
	w.workflow.connectionsFromTo[from] = to
	return w
}

func (w *WorkflowBuilder) Build() (*Workflow, error) {
	if err := w.workflow.Validate(); err != nil {
		return nil, err
	}
	return w.workflow, nil
}
