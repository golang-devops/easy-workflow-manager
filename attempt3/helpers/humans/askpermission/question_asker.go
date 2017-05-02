package askpermission

type QuestionAsker interface {
	GetAnswer() (string, error)
}
