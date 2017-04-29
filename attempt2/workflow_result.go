package attempt2

type WorkflowResult struct {
	Error             error
	LastAttemptedNode Node
	SuccessfulNodes   []Node
}
