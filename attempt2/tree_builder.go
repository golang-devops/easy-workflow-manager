package attempt2

type TreeBuilder interface {
	Build() (*Tree, error)
}

func NewTreeBuilder(initialNode Node) TreeBuilder {
	return &treeBuilder{
		tree: &Tree{
			self: initialNode,
		},
	}
}

type treeBuilder struct {
	tree *Tree
}

func (t *treeBuilder) Build() (*Tree, error) {
	return t.tree, nil
}
