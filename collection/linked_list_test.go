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
	L := collection.NewLinkedList[string]()
	assertion(t, 0, L.Len)
	L.Push("1")
	assertion(t, 1, L.Len)
}

func TestLinkedListCase2(t *testing.T) {
	L := collection.NewLinkedList("0")
	assertion(t, 1, L.Len)
	L.Push("1")
	assertion(t, 2, L.Len)
	I := L.IndexOf("1")
	assertion(t, 1, I)
}

func TestLinkedListCase3(t *testing.T) {
	U1 := User{
		Name:   "U1",
		Age:    20,
		Gendar: "male",
	}
	U2 := User{
		Name:   "U2",
		Age:    18,
		Gendar: "female",
	}
	L := collection.NewLinkedList(U1)
	assertion(t, 1, L.Len)
	L.Push(U2)
	assertion(t, 2, L.Len)
	I := L.IndexOf(U2)
	assertion(t, 1, I)
}

func TestLinkedListCase4(t *testing.T) {
	U1 := User{
		Name:   "U1",
		Age:    20,
		Gendar: "male",
	}
	U2 := User{
		Name:   "U2",
		Age:    18,
		Gendar: "female",
	}
	U3 := User{
		Name:   "U3",
		Age:    22,
		Gendar: "female",
	}
	U4 := User{
		Name:   "U4",
		Age:    24,
		Gendar: "male",
	}
	L := collection.NewLinkedList(U1, U2)
	assertion(t, 2, L.Len)
	L.Push(U3, U4)
	assertion(t, 4, L.Len)
	I := L.IndexOf(U3)
	assertion(t, 2, I)
	GetU3 := L.Get(I)
	assertion(t, U3, *GetU3)
}

func BenchmarkLinkedListCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		U1 := User{
			Name:   "U1",
			Age:    20,
			Gendar: "male",
		}
		U2 := User{
			Name:   "U2",
			Age:    18,
			Gendar: "female",
		}
		L := collection.NewLinkedList(U1)
		L.Push(U2)
		L.IndexOf(U2)
	}
}

func BenchmarkLinkedListCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		L := collection.NewLinkedList("u1")
		L.Push("u2")
		L.IndexOf("u2")
	}
}
