package benchmark

import (
	"fmt"
	"go-benchmark-demo/log"
	"reflect"

	"golang.org/x/exp/slices"
)

func SliceCreation() {
	slice := []string{"a1", "a2"}
	log.Debug(fmt.Sprintf("slice cap is %d", cap(slice)))
	slice = append(slice, "a3")
	log.Debug(fmt.Sprintf("slice cap is %d", cap(slice)))
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

func ArrayCreation() {
	array1 := [...]string{"a1", "a2"}
	array2 := [3]string{}
	array2[0] = array1[0]
	array2[1] = array1[1]
	log.PrintWithColor(
		log.Color.Blue,
		fmt.Sprintf(
			"array1 cap is [%d], array2 cap is [%d], array1 value is %+q, array2 value is %+q",
			cap(array1),
			cap(array2),
			array1,
			array2,
		),
	)
}

func ArrayCopy() {
	array1 := [...]string{"a1", "a2"}
	// A copy of an array by value, arr := arr1
	// A copy of an array by reference, arr := &arr1
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
