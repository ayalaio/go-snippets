package main

/* Interface used for Sort template */
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type Sortable struct{}

/* Sort template */
func Sort(a Interface) {
	for i := 0; i < a.Len()-1; i++ {
		for j := 0; j < a.Len()-i-1; j++ {

			if a.Less(j+1, j) {
				a.Swap(j, j+1)
			}
		}
	}
}

type Narco struct {
	Name string
	Age  int
}
type NarcoList []Narco

/* Following 3 funcs fill the Sort template */

func (nl NarcoList) Len() int {
	return len(nl)
}

func (nl NarcoList) Less(i, j int) bool {
	return nl[i].Age < nl[j].Age
}

func (nl NarcoList) Swap(i, j int) {
	nl[i], nl[j] = nl[j], nl[i]
}

func main() {
	pepe := Narco{Age: 12}
	lalo := Narco{Age: 34}
	rafa := Narco{Age: 17}

	narcos := NarcoList{}
	narcos = append(narcos, pepe, lalo, rafa)

	Sort(narcos)

	for i := 0; i < narcos.Len(); i++ {
		println(narcos[i].Age)
	}

}
