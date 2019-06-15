package main

import "fmt"

func binaryFind(a []int, target, s, f int) (int, error) {

	middle := s + (f-s)/2
	if a[middle] == target {
		return middle, nil
	} else if s == f {
		return -1, fmt.Errorf("Number not found")
	} else if target < a[middle] {
		return binaryFind(a, target, s, middle)
	} else if target > a[middle] {
		return binaryFind(a, target, middle+1, f)
	}

	return -1, fmt.Errorf("Impossible")
}

func findReps(a []int, target int) (int, error) {

	if len(a) == 0 {
		return -1, fmt.Errorf("Empty array is invalid")
	}

	pos, err := binaryFind(a, target, 0, len(a))
	if err == nil {
		i := pos
		j := pos
		for i >= 0 && a[i] == target {
			i--
		}
		for j < len(a) && a[j] == target {
			j++
		}
		i++
		j--
		return j - i + 1, nil
	}
	return pos, err
}

func main() {
	a := []int{1, 1, 1, 1, 1, 3, 3, 3, 3, 4, 4, 4, 4, 6, 6, 6, 6, 7, 7, 8, 8, 8, 9}

	numOfRepetitions, err := findReps(a, 2)
	if err == nil {
		fmt.Println(numOfRepetitions)
	} else {
		fmt.Println(err)
	}

	b := []int{}

	numOfRepetitions, err = findReps(b, 2)
	if err == nil {
		fmt.Println(numOfRepetitions)
	} else {
		fmt.Println(err)
	}

}
