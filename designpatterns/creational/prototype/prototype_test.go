package creational

import (
	"fmt"
	"testing"
)

func TestShirtDispenser(t *testing.T) {
	dispenser := ShirtDispenser{}

	var shirt Shirt

	shirt = dispenser.GetShirt(WHITE)
	whiteShirt := shirt.(*ShirtImpl)
	whiteShirt.SKU = 12324

	shirt = dispenser.GetShirt(WHITE)
	otherWhiteShirt := shirt.(*ShirtImpl)
	otherWhiteShirt.SKU = 12325

	fmt.Println(whiteShirt.GetInfo())
	fmt.Println(otherWhiteShirt.GetInfo())

	if whiteShirt.SKU == otherWhiteShirt.SKU {
		t.Error("Shirts should be unique")
	}
}
