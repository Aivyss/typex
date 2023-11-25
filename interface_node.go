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

type DefaultDescendNode[T any] struct {
	this        T
	descendants []DescendNode[T]
}

func (n *DefaultDescendNode[T]) AddDescendantNode(node DescendNode[T]) {
	n.descendants = append(n.descendants, node)
}

func (n *DefaultDescendNode[T]) This() T {
	return n.this
}

func (n *DefaultDescendNode[T]) HasDescendants() bool {
	return len(n.descendants) > 0
}

func (n *DefaultDescendNode[T]) AddDescendant(t T) {
	n.descendants = append(n.descendants, &DefaultDescendNode[T]{
		this:        t,
		descendants: []DescendNode[T]{},
	})
}

func (n *DefaultDescendNode[T]) GetDescendants() []DescendNode[T] {
	return n.descendants
}

func (n *DefaultDescendNode[T]) SetThis(t T) {
	n.this = t
}

func NewDescendNode[T any](t T) DescendNode[T] {
	return &DefaultDescendNode[T]{
		this:        t,
		descendants: []DescendNode[T]{},
	}
}

func NewMutableDescendNode[T any](t T) MutableDescendNode[T] {
	return &DefaultDescendNode[T]{
		this:        t,
		descendants: []DescendNode[T]{},
	}
}
