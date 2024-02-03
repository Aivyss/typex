package collection

type Set[E comparable] interface {
	Add(e E)
	AddAll(s []E)
	Contains(e E) bool
	ContainsAll(s []E) bool
	IsEmpty() bool
	Remove(e E)
	RemoveAll(s []E)
	ToSlice() []E
	Clear()
}

func NewSet[E comparable](elements ...E) Set[E] {
	obj := &hashSet[E]{
		internal: map[E]bool{},
	}

	obj.AddAll(elements)

	return obj
}

type hashSet[E comparable] struct {
	internal map[E]bool
}

func (h *hashSet[E]) Add(e E) {
	h.internal[e] = true
}

func (h *hashSet[E]) AddAll(s []E) {
	for _, element := range s {
		h.internal[element] = true
	}
}

func (h *hashSet[E]) Contains(e E) bool {
	exist, ok := h.internal[e]
	return exist && ok
}

func (h *hashSet[E]) ContainsAll(s []E) bool {
	checks := make([]bool, 0, len(s))
	for _, element := range s {
		checks = append(checks, h.Contains(element))
	}

	return Reduce(checks, func(prev bool, curr bool) bool {
		return prev && curr
	})
}

func (h *hashSet[E]) IsEmpty() bool {
	return len(h.ToSlice()) == 0
}

func (h *hashSet[E]) Remove(e E) {
	delete(h.internal, e)
}

func (h *hashSet[E]) RemoveAll(s []E) {
	for _, element := range s {
		h.Remove(element)
	}
}

func (h *hashSet[E]) ToSlice() []E {
	elements := make([]E, 0, len(h.internal))
	for element := range h.internal {
		if h.Contains(element) {
			elements = append(elements, element)
		}
	}

	return elements
}

func (h *hashSet[E]) Clear() {
	h.internal = map[E]bool{}
}
