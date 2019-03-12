package decorator

import "fmt"

type IngredientAdder interface {
	AddIngredient() (string, error)
}

type PizzaCore struct {
	Ingredient IngredientAdder
}

func (i *PizzaCore) AddIngredient() (string, error) {
	return "Pizza Core has Ingredients:", nil
}

type MeatIngredient struct {
	Ingredient IngredientAdder
}

func (i *MeatIngredient) AddIngredient() (string, error) {
	if i.Ingredient == nil {
		return "", fmt.Errorf("Cant have meat by itself.")
	}
	str, err := i.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s meat,", str), nil
}

type OnionIngredient struct {
	Ingredient IngredientAdder
}

func (i *OnionIngredient) AddIngredient() (string, error) {
	if i.Ingredient == nil {
		return "", fmt.Errorf("Cant have onion by itself.")
	}
	str, err := i.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s onion,", str), nil
}
