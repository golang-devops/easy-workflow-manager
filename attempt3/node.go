package attempt3

type Node interface {
	Name() string

	Accept(visitor NodeVisitor)
}

type NodeSlice []Node

func (n NodeSlice) ContainsNode(node Node) bool {
	for _, n2 := range n {
		if n2 == node {
			return true
		}
	}
	return false
}
