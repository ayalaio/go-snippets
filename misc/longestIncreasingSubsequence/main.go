package main

import "math"

func lengthOfLIS(nums []int) int {

	mem := make([][]int, len(nums)+1)
	for i := 0; i < len(mem); i++ {
		mem[i] = make([]int, len(nums)+1)
	}
	return LIS(nums, 0, -1, mem)
}

func LIS(nums []int, currPos, prevPos int, mem [][]int) int {

	if mem[currPos][prevPos+1] > 0 {
		return mem[currPos][prevPos+1]
	}

	if len(nums) == 0 {
		return 0
	}

	if currPos >= len(nums) {
		return 0
	}

	with := math.MinInt64

	// add it only if currnum is bigger than prevnum
	if prevPos == -1 || nums[currPos] > nums[prevPos] {
		with = 1 + LIS(nums, currPos+1, currPos, mem)
	}

	without := LIS(nums, currPos+1, prevPos, mem)

	mem[currPos][prevPos+1] = int(math.Max(float64(with), float64(without)))
	return mem[currPos][prevPos+1]

}

func main() {

}
