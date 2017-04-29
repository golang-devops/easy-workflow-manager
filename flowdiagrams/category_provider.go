package flowdiagrams

type CategoryProvider interface {
	Id() string
	Background() string
	TextColor() string
}

func NewSimpleCategoryProvider(id, background, textColor string) CategoryProvider {
	return &simpleCategoryProvider{
		id:         id,
		background: background,
		textColor:  textColor,
	}
}

type simpleCategoryProvider struct {
	id         string
	background string
	textColor  string
}

func (s *simpleCategoryProvider) Id() string         { return s.id }
func (s *simpleCategoryProvider) Background() string { return s.background }
func (s *simpleCategoryProvider) TextColor() string  { return s.textColor }
