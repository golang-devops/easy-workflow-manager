package attempt2

type Node interface {
	Execute(sharedData SharedData) error
}
