package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {

	mem := make(map[string]int)

	return coinChangeR(coins, amount, 0, 0, mem)

}

func coinChangeR(coins []int, amount, currentAmount, idx int, mem map[string]int) int {

	if currentAmount > amount {
		return -1
	}

	if idx >= len(coins) {
		return -1
	}

	h := fmt.Sprint(idx, ",", currentAmount)
	if v, ok := mem[h]; ok {
		return v
	}

	// not include the current coin
	notIncluded := -1
	resNotInc := coinChangeR(coins, amount, currentAmount, idx+1, mem)
	if resNotInc > -1 {
		notIncluded = resNotInc
	}

	// include the current coin
	included := -1
	if currentAmount == amount {
		//nothing else there to include
		included = 0
	} else {
		resInc := coinChangeR(coins, amount, currentAmount+coins[idx], idx, mem)
		if resInc > -1 {
			included = 1 + resInc
		}
	}

	better := -1
	if included > -1 && notIncluded > -1 {
		better = int(math.Min(float64(included), float64(notIncluded)))
	} else if included > -1 {
		better = included
	} else if notIncluded > -1 {
		better = notIncluded
	}

	mem[h] = better
	return better

}

func main() {}
