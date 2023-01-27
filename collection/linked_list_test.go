package collection_test

import (
	"go-benchmark-demo/collection"
	"testing"
)

func assertion(t *testing.T, assert any, result any) {
	if assert != result {
		t.Errorf("Error, %s, %s", result, assert)
	}
}

func TestLinkedList(t *testing.T) {
	l := collection.NewLinkedList[string]()
	assertion(t, 0, l.Len)
	l.Push("1")
	assertion(t, 1, l.Len)

	l = collection.NewLinkedList("0")
	assertion(t, 1, l.Len)
	l.Push("1")
	assertion(t, 2, l.Len)
}
