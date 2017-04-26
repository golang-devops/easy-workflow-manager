package elements

type Connection struct{}
type SequenceFlowConnection struct{ Connection }
type MessageFlowConnection struct{ Connection }
type AssociationConnection struct{ Connection }
