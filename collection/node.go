package collection

import (
	"bytes"
	"encoding/gob"
)

type Node[T interface{}] struct {
	Previous *Node[T]
	Next     *Node[T]
	Value    T
}

func (n *Node[T]) Hash() []byte {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(n)
	return b.Bytes()
}
