package decorator

import (
	"strings"
	"testing"
)

func TestDecorator(t *testing.T) {

	t.Run("Pizza Core can have no Ingredients", func(t *testing.T) {
		expectedMsg := "Pizza Core has Ingredients:"

		p := PizzaCore{}
		str, err := p.AddIngredient()
		if err != nil {
			t.Fatal(err)
		}
		if str != expectedMsg {
			t.Error("Pizza core has no expected msg")
		}
	})

	t.Run("Meat can't be by itself", func(t *testing.T) {
		m := MeatIngredient{}
		_, err := m.AddIngredient()
		if err == nil {
			t.Error("We expect an error of a pizzaless meat")
		}
	})

	t.Run("Meat is an Ingredient", func(t *testing.T) {
		expectedMsg := "meat"

		m := MeatIngredient{&PizzaCore{}}
		str, err := m.AddIngredient()
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(str, expectedMsg) {
			t.Error("Meat has no expected msg")
		}
	})

	t.Run("Meat + Onion is an Ingredient", func(t *testing.T) {
		expectedMsg := "meat, onion"

		o := &OnionIngredient{
			&MeatIngredient{
				&PizzaCore{},
			},
		}
		str, err := o.AddIngredient()
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(str, expectedMsg) {
			t.Error("Meat, Onion has no expected msg")
		}
	})

}
