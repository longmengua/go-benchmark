package collection

type Set[T interface{}] struct {
	Len    int
	Values map[interface{}]bool
}

func (s *Set[T]) Push(n Node[T]) *Set[T] {
	if !s.Values[n.Hash()] {
		s.Values[n.Hash()] = true
	}
	return s
}
