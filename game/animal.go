package game

import (
	"fmt"
)

/**** Dog ****/
type dog struct {
	name  string
	speed int
	x     int
}

func NewDog(name string) *dog {
	return &dog{name, 2, 0}
}

func (d *dog) Move() {
	d.x += d.speed
}

func (d *dog) Name() string {
	return d.name
}

func (d *dog) Position() int {
	return d.x
}

/**** Cat ****/
type cat struct {
	name  string
	speed int
	x     int
}

func NewCat(name string) *cat {
	return &cat{name, 5, 0}
}

func (c *cat) Name() string {
	return c.name
}

func (c *cat) Position() int {
	return c.x
}

func (c *cat) Move() {
	c.x += c.speed
}

func (c *cat) String() string {
	return c.name + " is winner!"
}

/**** Owl ****/
type owl struct {
	name  string
	speed int
	x     int
}

func NewOwl(name string) *owl {
	return &owl{name, 3, 0}
}

func (o *owl) Move() {
	o.x += o.speed
}

func (o *owl) Name() string {
	return o.name
}

func (o *owl) Position() int {
	return o.x
}

func (o *owl) String() string {
	return fmt.Sprintf("Owl %s at %d", o.name, o.x)
}
