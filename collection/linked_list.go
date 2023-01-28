package collection

import "fmt"

type LinkedList[T interface{}] struct {
	len      int
	isSorted bool
	first    *Node[T]
	last     *Node[T]
}

func NewLinkedList[T interface{}](values ...T) *LinkedList[T] {
	l := LinkedList[T]{
		len:      0,
		isSorted: false,
		first:    nil,
		last:     nil,
	}
	for _, value := range values {
		l.push(value)
	}
	return &l
}

func (l *LinkedList[T]) push(value T) *LinkedList[T] {
	n := Node[T]{
		value:    value,
		previous: nil,
		next:     nil,
	}
	if l.len == 0 {
		l.first = &n
		l.last = &n
	} else {
		l.last.next = &n
		n.previous = l.last
		l.last = &n
		n.next = l.first
	}
	l.len++
	return l
}

func (l *LinkedList[T]) indexOf(value T) int {
	if l.len > 0 {
		n := l.first
		for i := 0; i < l.len; i++ {
			v1 := fmt.Sprintf("%v", n.value)
			v2 := fmt.Sprintf("%v", value)
			if v1 == v2 {
				return i
			}
			n = n.next
		}
	}
	return -1
}

func (l *LinkedList[T]) Len(values ...T) int {
	return l.len
}

func (l *LinkedList[T]) First() T {
	return l.first.value
}

func (l *LinkedList[T]) Last() T {
	return l.last.value
}

func (l *LinkedList[T]) Push(values ...T) *LinkedList[T] {
	for _, v := range values {
		l.push(v)
	}
	return l
}

func (l *LinkedList[T]) Pop() T {
	var n *Node[T]
	if l.len > 0 {
		n = l.last
		l.last = l.last.previous
		l.last.next = l.first
		l.len--
	}
	return n.value
}

func (l *LinkedList[T]) Get(index int) *T {
	halfIndex := l.len / 2
	var node *Node[T]
	if index < halfIndex {
		node = l.first
		for i := 0; i < index; i++ {
			if i == index {
				return &node.value
			}
			node = node.next
		}
	}

	node = l.last
	for i := 0; i < index; i++ {
		if i+1 == index {
			return &node.value
		}
		node = node.previous
	}
	return nil
}

func (l *LinkedList[T]) IndexOf(value T) int {
	return l.indexOf(value)
}

func (l *LinkedList[T]) Contains(value T) bool {
	return l.indexOf(value) > -1
}

func (l *LinkedList[T]) Values() []T {
	values := make([]T, l.len)
	node := l.first
	for i := 0; node != l.last; i++ {
		values[i] = node.value
		node = node.next
	}
	return values
}
