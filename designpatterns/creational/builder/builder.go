package creational

type BuilderProcess interface {
	SetWheels() BuilderProcess
	SetSeats() BuilderProcess
	SetStructure() BuilderProcess
	GetVehicle() Vehicle
}

//Director
type ManufacturingDirector struct {
	builder BuilderProcess
}

func (m *ManufacturingDirector) SetBuilder(b BuilderProcess) {
	m.builder = b
}

func (m *ManufacturingDirector) Construct() {
	m.builder.SetWheels().SetSeats().SetStructure()
}

// Product
type Vehicle struct {
	Wheels    int
	Seats     int
	Structure string
}

// Builder
type CarBuilder struct {
	v Vehicle
}

func (c *CarBuilder) SetWheels() BuilderProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuilderProcess {
	c.v.Seats = 4
	return c
}

func (c *CarBuilder) SetStructure() BuilderProcess {
	c.v.Structure = "This is a car"
	return c
}

func (c *CarBuilder) GetVehicle() Vehicle {
	return c.v
}
