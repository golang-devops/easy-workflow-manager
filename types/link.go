package types

type Link interface {
	From() Node
	To() Node
}
