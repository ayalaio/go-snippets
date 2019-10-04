package main

import "fmt"

func mergeSort(a []int) {
	mergeSortR(a, 0, len(a)-1)
}

func mergeSortR(a []int, init, final int) {
	if init >= final {
		return
	}

	if init+1 == final {
		if a[init] > a[final] {
			a[init], a[final] = a[final], a[init]
		}
		return
	}

	middle := (init + final) / 2

	mergeSortR(a, init, middle)
	mergeSortR(a, middle+1, final)

	b := make([]int, final-init+1)

	i := init
	j := middle + 1
	k := 0

	for i <= middle && j <= final {
		if a[i] <= a[j] {
			b[k] = a[i]
			i++
			k++
		} else if a[j] <= a[i] {
			b[k] = a[j]
			j++
			k++
		}
	}

	for i <= middle {
		b[k] = a[i]
		i++
		k++
	}

	for j <= final {
		b[k] = a[j]
		j++
		k++
	}

	for i, j := init, 0; i <= final; i, j = i+1, j+1 {
		a[i] = b[j]
	}

}

func main() {
	a := []int{5, 3, 8, 1, 9, 0, 5, 4, 4, 7, 8, 9, 0, 2, 4, 5}

	//a := []int{9, 8, 7, 6, 0, 4, 3, 5, 1}

	mergeSort(a)

	fmt.Println(a)

}
