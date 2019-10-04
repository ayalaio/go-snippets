package main

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Live() {
	println("I'm alive")
}

type Eater interface {
	Eat()
}

type Human struct {
	*Animal
	Eater
	Friend         Animal
	CustomAnoyment func()
}

type EaterImplForHumans struct {
}

func (e EaterImplForHumans) Eat() {
	println("chew")
}

func main() {
	var h Human
	h = Human{
		Animal: &Animal{Name: "Pablo"},
		Eater:  &EaterImplForHumans{}, // This is composite pattern
		Friend: Animal{Name: "Firulais"},
		CustomAnoyment: func(h *Human) func() {
			return func() {
				fmt.Printf("%s anoys %s with Pokemon\n", h.Name, h.Friend.Name)
			}
		}(&h),
	}
	println(h.Name)
	h.Eat()
	h.CustomAnoyment()
	h.Live()
}
