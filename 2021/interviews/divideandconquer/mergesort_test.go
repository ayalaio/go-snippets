package divideandconquer

import (
	"fmt"
	"reflect"
	"testing"
)

func MergeSort(array []int) {

	if len(array) == 0 || len(array) == 1 {
		return
	}

	if len(array) == 2 {
		if array[0] > array[1] {
			array[0], array[1] = array[1], array[0]
		}
		return
	}

	l := len(array)
	l_half := l / 2
	arrayLeft := array[0:l_half]
	arrayRight := array[l_half:l]
	MergeSort(arrayLeft)
	MergeSort(arrayRight)

	iL, iR := 0, 0
	arrayCopy := make([]int, len(array))
	for i := 0; i < len(arrayCopy); i++ {
		do := ""
		if iL >= len(arrayLeft) {
			do = "right"
		} else if iR >= len(arrayRight) {
			do = "left"
		} else if arrayRight[iR] < arrayLeft[iL] {
			do = "right"
		} else if arrayLeft[iL] <= arrayRight[iR] {
			do = "left"
		} else {
			break
		}
		if do == "right" {
			arrayCopy[i] = arrayRight[iR]
			iR++
		}
		if do == "left" {
			arrayCopy[i] = arrayLeft[iL]
			iL++
		}
	}
	for i := 0; i < len(array); i++ {
		array[i] = arrayCopy[i]
	}
}

func Test_MergeSort(t *testing.T) {
	ary := []int{1, 9, 2, 8, 3, 7, 4, 5, 6, 9, 2, 7, 3, 0}
	expectedResult := []int{0, 1, 2, 2, 3, 3, 4, 5, 6, 7, 7, 8, 9, 9}
	MergeSort(ary)
	fmt.Println(ary)
	if !reflect.DeepEqual(ary, expectedResult) {
		t.Error("Array is not sorted")
	}
}
