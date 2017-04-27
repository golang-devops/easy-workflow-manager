package types

type WorkflowResult struct {
	Error             error
	LastAttemptedNode Node
	SuccessfulNodes   []Node
}
