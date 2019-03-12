package creational

import (
	"errors"
	"fmt"
)

const (
	HUMAN = 1
	DOG   = 2
)

type Human struct {
	Feet int
}

type Animal interface {
	Heal()
}

func (c *Human) Heal() {
	c.Feet = 2
}

func GetBrokenAnimal(animalType int) (animal Animal, err error) {
	switch animalType {
	case HUMAN:
		animal, err = &Human{Feet: 1}, nil
	case DOG:
		animal, err = nil, errors.New("NO DOG IMPLEMENTATION")
	}
	return
}

func main() {

	animal, _ := GetBrokenAnimal(HUMAN)
	animal.Heal()

	human, ok := animal.(*Human)

	fmt.Println(human.Feet, ok)

}
