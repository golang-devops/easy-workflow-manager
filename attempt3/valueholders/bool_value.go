package valueholders

type BoolValue interface {
	Set(value bool)
	Get() bool
}

func NewBoolValue() BoolValue {
	return &boolValue{}
}

type boolValue struct {
	val bool
}

func (s *boolValue) Set(value bool) { s.val = value }
func (s *boolValue) Get() bool      { return s.val }
