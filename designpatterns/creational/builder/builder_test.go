package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturer := new(ManufacturingDirector)
	carBuilder := new(CarBuilder)
	manufacturer.SetBuilder(carBuilder)
	manufacturer.Construct()
	carVehicle := carBuilder.GetVehicle()

	if carVehicle.Wheels != 4 {
		t.Error("Wheels should be 0")
	}

	if carVehicle.Seats != 4 {
		t.Error("Seats should be 4")
	}

	if carVehicle.Structure != "This is a car" {
		t.Error("Wrong Structure")
	}

}
