package types

type Condition interface {
	Node

	ConditionAnswerProvider() ConditionAnswerProvider

	WhenTrue() Node
	WhenFalse() Node
}

type ConditionAnswerProvider interface {
	ConditionAnswer() bool
}
