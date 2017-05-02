package valueholders

type StringValue interface {
	Set(value string)
	Get() string
}

func NewStringValue() StringValue {
	return &stringValue{}
}

type stringValue struct {
	val string
}

func (s *stringValue) Set(value string) { s.val = value }
func (s *stringValue) Get() string      { return s.val }
