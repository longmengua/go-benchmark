package collection

type LinkedList[T interface{}] struct {
	Len  int
	Head *Node[T]
	Tail *Node[T]
}

func NewLinkedList[T interface{}](values ...T) *LinkedList[T] {
	l := LinkedList[T]{
		Len:  0,
		Head: nil,
		Tail: nil,
	}
	for _, value := range values {
		l.Push(value)
	}
	return &l
}

func (l *LinkedList[T]) Push(value T) *LinkedList[T] {
	n := Node[T]{
		Value:    value,
		Previous: nil,
		Next:     nil,
	}
	if l.Len == 0 {
		l.Head = &n
		l.Tail = &n
	} else {
		l.Tail.Next = &n
		l.Tail = &n
		n.Next = l.Head
	}
	l.Len++
	return l
}

func (l *LinkedList[T]) Pop() *Node[T] {
	var n *Node[T]
	if l.Len > 0 {
		n = l.Tail
		l.Tail = l.Tail.Previous
		l.Tail.Next = l.Head
		l.Len--
	}
	return n
}

func (l *LinkedList[T]) IndexOf(node *Node[T]) int {
	if l.Len > 0 {
		n := l.Head
		for i := 0; i < l.Len; i++ {
			if n == node {
				return i
			}
			n = l.Head.Next
		}
	}
	return -1
}

func (l *LinkedList[T]) Contains(node *Node[T]) bool {
	return l.IndexOf(node) > -1
}

func (l *LinkedList[T]) Value() []T {
	values := make([]T, l.Len)
	node := l.Head
	for i := 0; node != l.Tail; i++ {
		values[i] = node.Value
		node = node.Next
	}
	return values
}
