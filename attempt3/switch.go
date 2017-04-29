package attempt3

type Switch interface {
	Node

	SwitchAnswerProvider() SwitchAnswerProvider
}

type SwitchAnswerProvider interface {
	AllCases() []Node
	SwitchAnswer() Node
}
