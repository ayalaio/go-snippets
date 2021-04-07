package main

import "fmt"

type Product struct {
	name          string
	value, weight int
}

var mem [][]int

func areProductsLeft(products []*Product, prodIdx, capacityLeft int) bool {
	for i := prodIdx; i < len(products); i++ {
		if products[i].weight <= capacityLeft {
			return true
		}
	}
	return false
}

func FillTheBag(capacity int, products []*Product, prodIdx, currVal, currWeight int) int {

	if currWeight > capacity {
		return 0
	}

	if prodIdx >= len(products) {
		return 0
	}

	if v := mem[prodIdx][currWeight]; v > 0 {
		return v
	}

	if currWeight <= capacity && !areProductsLeft(products, prodIdx, capacity-currWeight) {
		mem[prodIdx][currWeight] = currVal
		return currVal
	}

	// We include product
	withOut := FillTheBag(capacity, products, prodIdx+1,
		currVal+products[prodIdx].value, currWeight+products[prodIdx].weight)

	// We don't include product
	with := FillTheBag(capacity, products, prodIdx+1,
		currVal, currWeight)

	max := maxInt(withOut, with)
	mem[prodIdx][currWeight] = max
	return max

}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	totalCapacity := 165

	products := []*Product{
		&Product{"A", 92, 23},
		&Product{"B", 57, 31},
		&Product{"C", 49, 29},
		&Product{"D", 68, 44},
		&Product{"E", 60, 53},
		&Product{"F", 43, 38},
		&Product{"G", 67, 63},
		&Product{"H", 84, 85},
		&Product{"I", 87, 89},
		&Product{"J", 72, 82},
	}

	mem = make([][]int, len(products))
	for i := range mem {
		mem[i] = make([]int, totalCapacity+1)
	}

	maxValue := FillTheBag(totalCapacity, products, 0, 0, 0)

	fmt.Println(maxValue)

}
