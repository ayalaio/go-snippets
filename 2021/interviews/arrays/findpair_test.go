package arrays

import (
	"errors"
	"testing"
)

func FindPairGivenSum(ary []int, target int) (int, int, error) {
	inStock := make(map[int]bool)
	for _, n := range ary {
		if _, ok := inStock[target-n]; ok {
			return n, target - n, nil
		}
		inStock[n] = true
	}
	return 0, 0, errors.New("No pair")
}

func Test_FindPairGivenSum_WhenPairExists(t *testing.T) {
	target := 9
	sampleArray := []int{1, 2, 3, 4, 5, 6, 7}
	a, b, _ := FindPairGivenSum(sampleArray, target)
	if a+b != target {
		t.Error("Function is not returning values that equal the sum")
	}
}

func Test_FindPairGivenSum_WhenNoPairExists(t *testing.T) {
	target := 55
	sampleArray := []int{1, 2, 3, 4, 5, 6, 7}
	_, _, err := FindPairGivenSum(sampleArray, target)
	if err == nil {
		t.Error("Function should have returned error")
	}
}
