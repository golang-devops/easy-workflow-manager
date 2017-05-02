package attempt3

func NewSimpleSwitch(name string, cases []*SwitchCase) Switch {
	return &simpleSwitch{
		name:  name,
		cases: cases,
	}
}

type simpleSwitch struct {
	name  string
	cases []*SwitchCase
}

func (s *simpleSwitch) Name() string               { return s.name }
func (s *simpleSwitch) Cases() []*SwitchCase       { return s.cases }
func (s *simpleSwitch) Accept(visitor NodeVisitor) { visitor.VisitSwitch(s) }
