package types

func NewSimpleActivity(name string, executor func() error, nextNode Node) Activity {
	return &simpleActivity{
		name:     name,
		executor: executor,
		nextNode: nextNode,
	}
}

type simpleActivity struct {
	name     string
	executor func() error
	nextNode Node
}

func (s *simpleActivity) Name() string               { return s.name }
func (s *simpleActivity) Accept(visitor NodeVisitor) { visitor.VisitActivity(s) }
func (s *simpleActivity) Execute() error             { return s.executor() }
func (s *simpleActivity) Next() Node                 { return s.nextNode }
