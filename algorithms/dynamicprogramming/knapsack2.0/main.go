package main

import "fmt"

type Product struct {
	name          string
	value, weight int
}

func areProductsLeft(products []*Product, prodIdx, capacityLeft int) bool {
	for i := prodIdx; i < len(products); i++ {
		if products[i].weight <= capacityLeft {
			return true
		}
	}
	return false
}

func FillTheBag(capacity int, products []*Product, prodIdx, currVal, currWeight int) int {

	if currWeight <= capacity && !areProductsLeft(products, prodIdx, capacity-currWeight) {
		return currVal
	}

	if currWeight > capacity {
		return 0
	}

	if prodIdx >= len(products) {
		return 0
	}

	// We include product
	withOut := FillTheBag(capacity, products, prodIdx+1,
		currVal+products[prodIdx].value, currWeight+products[prodIdx].weight)

	// We don't include product
	with := FillTheBag(capacity, products, prodIdx+1,
		currVal, currWeight)

	return maxInt(withOut, with)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
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

	totalCapacity := 165

	maxValue := FillTheBag(totalCapacity, products, 0, 0, 0)

	fmt.Println(maxValue)

}
