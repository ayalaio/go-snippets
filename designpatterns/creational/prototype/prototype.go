package creational

import "fmt"

// White shirt prototype
// Blie shirt prototype
// Red shirt prototype

const (
	WHITE = 1
	BLUE  = 2
	RED   = 3
)

type Shirt interface {
	GetInfo() string
}

type ShirtImpl struct {
	Color int
	SKU   int
}

func (shirt ShirtImpl) GetInfo() string {
	return fmt.Sprintf("Im a COLORED %d shirt with SKU %d", shirt.Color, shirt.SKU)
}

type ShirtDispenser struct{}

func (_ ShirtDispenser) GetShirt(shirtType int) (shirt Shirt) {
	switch shirtType {
	case WHITE:
		newItem := *whiteShirt
		shirt = &newItem
	case BLUE:
		newItem := *blueShirt
		shirt = &newItem
	case RED:
		newItem := *redShirt
		shirt = &newItem
	}
	return
}

var whiteShirt *ShirtImpl = &ShirtImpl{
	Color: WHITE,
	SKU:   0,
}

var blueShirt *ShirtImpl = &ShirtImpl{
	Color: BLUE,
	SKU:   0,
}

var redShirt *ShirtImpl = &ShirtImpl{
	Color: RED,
	SKU:   0,
}
