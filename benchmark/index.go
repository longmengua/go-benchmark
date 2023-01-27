package benchmark

import (
	"fmt"
	"go-benchmark-demo/log"
	"reflect"

	"golang.org/x/exp/slices"
)

func SliceCreation() {
	slice := []string{"a1", "a2"}
	log.Debug(fmt.Sprintf("1. slice cap is %d", cap(slice)))
	slice = append(slice, "a3")
	log.Debug(fmt.Sprintf("2. slice cap is %d", cap(slice)))
}

func SliceCopy() {
	slice1 := []string{"a1", "a2"}
	slice2 := slice1
	log.Debug(fmt.Sprintf(
		"array1: [%p], array2: [%p], address is same: [%t], value is same [%t], (reflect) value is same [%t]",
		&slice1,
		&slice2,
		&slice1 == &slice2,
		slices.Equal(slice1, slice2),
		reflect.DeepEqual(slice1, slice2),
	))
}

func ArrayCopy() {
	array1 := [...]string{"a1", "a2"}
	array2 := array1
	log.Debug(fmt.Sprintf(
		"array1: [%p], array2: [%p], address is same: [%t], value is same [%t], (reflect) value is same [%t]",
		&array1,
		&array2,
		&array1 == &array2,
		array1 == array2,
		reflect.DeepEqual(array1, array2),
	))
}
