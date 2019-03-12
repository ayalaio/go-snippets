package main

type Visitor interface {
	Visit(Product)
}

type Visitable interface {
	Accept(Visitor)
}

type Product struct {
	Price int
}

func (p *Product) GetPrice() int {
	return p.Price
}

func (p *Product) Accept(v Visitor) {
	v.Visit(*p)
}

type ProductStock struct {
	Total int
}

func (ps *ProductStock) Visit(p Product) {
	ps.Total = ps.Total + p.GetPrice()
}

type Strawberry struct {
	Product
}
type Banana struct {
	Product
}
type Mango struct {
	Product
}

func main() {
	products := []Visitable{
		&Strawberry{
			Product: Product{
				Price: 1,
			},
		},
		&Banana{
			Product: Product{
				Price: 2,
			},
		},
		&Mango{
			Product: Product{
				Price: 3,
			},
		},
	}

	ps := ProductStock{}

	for _, p := range products {
		p.Accept(&ps)
	}

	println(ps.Total)
}
