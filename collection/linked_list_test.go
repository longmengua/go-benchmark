package collection_test

import (
	"go-benchmark-demo/collection"
	"testing"
)

type User struct {
	Name   string
	Age    int
	Gendar string
}

func assertion(t *testing.T, assert any, result any) {
	if assert != result {
		t.Errorf("Error, %s, %s", result, assert)
	}
}

func TestLinkedListCase1(t *testing.T) {
	// case 1.
	l := collection.NewLinkedList[string]()
	assertion(t, 0, l.Len)
	l.Push("1")
	assertion(t, 1, l.Len)
}

func TestLinkedListCase2(t *testing.T) {
	// case 2.
	l := collection.NewLinkedList("0")
	assertion(t, 1, l.Len)
	l.Push("1")
	assertion(t, 2, l.Len)
	i := l.IndexOf("1")
	assertion(t, 1, i)
}

func TestLinkedListCase3(t *testing.T) {
	// case 3.
	u1 := User{
		Name:   "U1",
		Age:    20,
		Gendar: "male",
	}
	u2 := User{
		Name:   "U2",
		Age:    18,
		Gendar: "female",
	}
	l1 := collection.NewLinkedList(u1)
	assertion(t, 1, l1.Len)
	l1.Push(u2)
	assertion(t, 2, l1.Len)
	i := l1.IndexOf(u2)
	assertion(t, 1, i)
}

func BenchmarkLinkedListCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u1 := User{
			Name:   "U1",
			Age:    20,
			Gendar: "male",
		}
		u2 := User{
			Name:   "U2",
			Age:    18,
			Gendar: "female",
		}
		l1 := collection.NewLinkedList(u1)
		l1.Push(u2)
		l1.IndexOf(u2)
	}
}

func BenchmarkLinkedListCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l1 := collection.NewLinkedList("u1")
		l1.Push("u2")
		l1.IndexOf("u2")
	}
}
