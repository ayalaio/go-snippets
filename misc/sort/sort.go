package main

import "fmt"

func bubblesort(ary *[]int) {
	a := *ary
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func mergesort(ary *[]int, initial, final int) {
	a := *ary
	if initial == final {
		return
	}

	var middle int
	middle = (initial + final) / 2

	mergesort(ary, initial, middle)
	mergesort(ary, middle+1, final)

	tmp := make([]int, final-initial+1)
	posAry1 := initial
	posAry2 := middle + 1

	for k := 0; k < len(tmp); k++ {
		if posAry2 > final {
			tmp[k] = a[posAry1]
			posAry1++
			continue
		}
		if posAry1 > middle {
			tmp[k] = a[posAry2]
			posAry2++
			continue
		}
		if a[posAry1] <= a[posAry2] {
			tmp[k] = a[posAry1]
			posAry1++
		} else {
			tmp[k] = a[posAry2]
			posAry2++
		}
	}
	for l := 0; l < len(tmp); l++ {
		a[l+initial] = tmp[l]
	}
}

func main() {
	a := []int{6, 43, 7, 89, 32, 1, 89, 3, 8, 3, 8, 1}

	bubblesort(&a)

	fmt.Println(a)

	b := []int{6, 43, 7, 89, 32, 1, 89, 3, 8, 3, 8, 1}

	mergesort(&b, 0, len(b)-1)

	fmt.Println(b)
}
