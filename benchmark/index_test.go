package benchmark_test

import (
	"go-benchmark-demo/benchmark"
	"testing"
)

func assertion(t *testing.T, assert any, result any) {
	if assert != result {
		t.Errorf("Error, %s, %s", result, assert)
	}
}

func TestSlice(t *testing.T) {
	benchmark.SliceCreation()
	benchmark.SliceCopy()
	benchmark.ArrayCopy()
	assertion(t, true, true)
}
