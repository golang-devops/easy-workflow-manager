package types

type Node interface {
	Execute(sharedData SharedData) error
}
