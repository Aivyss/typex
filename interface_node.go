package typex

type DescendNode[T any] interface {
	HasDescendants() bool
	AddDescendant(t T)
	AddDescendantNode(node DescendNode[T])
	GetDescendants() []DescendNode[T]
	This() T
}

type MutableDescendNode[T any] interface {
	DescendNode[T]
	SetThis(t T)
}

type defaultDescendNode[T any] struct {
	this        T
	descendants []DescendNode[T]
}

func (n *defaultDescendNode[T]) AddDescendantNode(node DescendNode[T]) {
	n.descendants = append(n.descendants, node)
}

func (n *defaultDescendNode[T]) This() T {
	return n.this
}

func (n *defaultDescendNode[T]) HasDescendants() bool {
	return len(n.descendants) > 0
}

func (n *defaultDescendNode[T]) AddDescendant(t T) {
	n.descendants = append(n.descendants, &defaultDescendNode[T]{
		this:        t,
		descendants: []DescendNode[T]{},
	})
}

func (n *defaultDescendNode[T]) GetDescendants() []DescendNode[T] {
	return n.descendants
}

func (n *defaultDescendNode[T]) SetThis(t T) {
	n.this = t
}

func NewDescendNode[T any](t T) DescendNode[T] {
	return &defaultDescendNode[T]{
		this:        t,
		descendants: []DescendNode[T]{},
	}
}

func NewMutableDescendNode[T any](t T) MutableDescendNode[T] {
	return &defaultDescendNode[T]{
		this:        t,
		descendants: []DescendNode[T]{},
	}
}
