package main

import "fmt"

type MemoizeBag struct {
	weight, value int
	products      []*Product
}

type Product struct {
	name   string
	value  int
	weight int
}

type ProductList []*Product

func (self ProductList) String() string {
	s := ""
	for i := 0; i < len(self); i++ {
		s += self[i].name
	}
	return s
}

func (self Product) String() string {
	return fmt.Sprintf("%s@%d,%d", self.name, self.value, self.weight)
}

var memo map[string]*MemoizeBag = make(map[string]*MemoizeBag)

func fillTheBag(totalCapacity int, products []*Product) (int, int, []*Product) {

	var mm ProductList
	mm = ProductList(products)
	if v, ok := memo[mm.String()]; ok {
		return v.weight, v.value, v.products
	}

	currWeight := 0
	currValue := 0
	for i := 0; i < len(products); i++ {
		currWeight += products[i].weight
		currValue += products[i].value
	}

	maxValue := 0
	maxWeight := 0
	maxProducts := []*Product{}

	if currWeight <= totalCapacity {
		maxWeight = currWeight
		maxValue = currValue
		maxProducts = products
	}

	for i := 0; i < len(products); i++ {
		a, b := products[0:i], products[i+1:]
		c := []*Product{}
		c = append(c, a...)
		c = append(c, b...)
		selectedWeight, selectedValue, selectedProducts := fillTheBag(totalCapacity, c)
		if selectedWeight <= totalCapacity && selectedValue > maxValue {
			maxWeight, maxValue, maxProducts = selectedWeight, selectedValue, selectedProducts
		}
	}

	memo[mm.String()] = &MemoizeBag{maxWeight, maxValue, maxProducts}

	return maxWeight, maxValue, maxProducts

}

func FillTheBag(totalCapacity int, products []*Product) (int, int, []*Product) {
	return fillTheBag(totalCapacity, products)
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

	sW, sV, sP := FillTheBag(totalCapacity, products)

	fmt.Println(sW, sV, sP)

}
