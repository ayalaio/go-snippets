package main

import (
	"fmt"
	"math"
)

func fillTheBackPack(maxCapacity int, products [][]int) int {
	m := [][]int{}

	topMax := 0

	for i := 0; i < len(products)+1; i++ { // + 1 row for no products
		m = append(m, make([]int, maxCapacity+1)) // + 1 for the 0-maxCapacity capacity
	}

	for i := 1; i < len(m); i++ { // start in row 1, as row 0 (no products) should be full of 0's
		for cap := 0; cap < len(m[0]); cap++ {
			productValue, productWeight := products[i-1][0], products[i-1][1]

			valWithoutThisProduct := m[i-1][cap]

			if cap-productWeight >= 0 { // Fits inside the backpack
				valWithThisProduct := productValue + m[i-1][cap-productWeight]
				m[i][cap] = int(math.Max(float64(valWithoutThisProduct), float64(valWithThisProduct)))
			} else {
				m[i][cap] = valWithoutThisProduct
			}

			if topMax < m[i][cap] {
				topMax = m[i][cap]
			}
		}
	}

	return topMax
}

func main() {
	products := [][]int{
		// V, W
		{92, 23},
		{57, 31},
		{49, 29},
		{68, 44},
		{60, 53},
		{43, 38},
		{67, 63},
		{84, 85},
		{87, 89},
		{72, 82},
	}
	totalCapacity := 165
	totalValue := fillTheBackPack(totalCapacity, products)
	fmt.Println(totalValue)
}
