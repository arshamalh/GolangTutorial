package car

import "fmt"

// Methods are used when we want to define specific functionality for custom types

type FuelTank struct {
	Capacity int
	Filled   int
}

type Car struct {
	Producer string
	Model    string
	Color    string
	Power    int
	IsOn     bool
	FuelTank
}

func main() {
	bmw_i8 := Car{
		Producer: "BMW",
		Model:    "i8",
		Color:    "White",
	}
	mc_720s := Car{
		Producer: "McLaren",
		Model:    "720 S",
		Color:    "Black",
	}

	bmw_i8.IsOn = true
	bmw_i8.Drive()

	fmt.Println("-------------")
	mc_720s.Drive()

}

func (f *FuelTank) Fill() {
	f.Filled = f.Capacity
}

func (c Car) Drive() Car {
	name := c.Producer + " - " + c.Model

	if c.IsOn {
		fmt.Println(name, "car is driving")
		c.FuelTank.Filled -= 10
		fmt.Println(c.FuelTank.Filled)
	} else {
		fmt.Println(name, "can't be driven without getting on")
	}
	return c
}

func (c *Car) Start() {
	c.IsOn = true
}

func (c *Car) FillTank() {
	c.FuelTank.Fill()
}
