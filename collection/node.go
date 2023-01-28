package collection

import (
	"bytes"
	"encoding/gob"
)

type Node[T interface{}] struct {
	previous *Node[T]
	next     *Node[T]
	value    T
}

func (n *Node[T]) Value() T {
	return n.value
}

func (n *Node[T]) Hash() []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(n)
	return b.Bytes()
}
