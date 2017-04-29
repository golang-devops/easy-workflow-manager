package attempt2

// type WorkflowExecutor interface {
// 	Execute(sharedData SharedData) *WorkflowResult
// }

// func newWorkflowExecutor(tree *Tree, sharedData SharedData) WorkflowExecutor {
// 	return &workflowExecutor{
// 		tree:       tree,
// 		sharedData: sharedData,
// 	}
// }

// type workflowExecutor struct {
// 	tree       *Tree
// 	sharedData SharedData
// }

// func (w *workflowExecutor) Execute() (result *WorkflowResult) {
// 	result = &WorkflowResult{}

// 	currentTree := w.tree
// 	for {
// 		node := currentTree.self
// 		result.LastAttemptedNode = node

// 		if err := node.Execute(w.sharedData); err != nil {
// 			result.Error = err
// 			return
// 		}
// 		result.SuccessfulNodes = append(result.SuccessfulNodes, node)

// 		currentTree = currentTree.decider.Decide(w.sharedData, currentTree.children)
// 	}
// }

// type workflow struct {
// 	startingNode Node
// 	nodes        []Node
// 	links        []Link
// }

// func (w *workflow) Execute(sharedData SharedData) (result *WorkflowResult) {
// 	result = &WorkflowResult{}

// 	result.LastAttemptedNode = w.startingNode
// 	if err := w.startingNode.Execute(sharedData); err != nil {
// 		result.Error = err
// 		return
// 	}

// 	currentTree := w.tree
// 	for {
// 		node := currentTree.self
// 		result.LastAttemptedNode = node

// 		if err := node.Execute(w.sharedData); err != nil {
// 			result.Error = err
// 			return
// 		}
// 		result.SuccessfulNodes = append(result.SuccessfulNodes, node)

// 		currentTree = currentTree.decider.Decide(w.sharedData, currentTree.children)
// 	}
// }
