package car

import "fmt"

type Car struct {
	year int
	name string
}

type Vehicle interface {
	setName(string)
	setYear(int)
	getName() string
	getYear() int
	print()
}

func (c *Car) setName(name string) {
	c.name = name
}

func (c *Car) setYear(year int) {
	c.year = year
}

func (c Car) getName() string {
	return c.name
}

func (c Car) getYear() int {
	return c.year
}
func (c Car) print() {
	fmt.Println(c.getName(), c.getYear())
}

func Initialize() {
	var a Car = Car{2018, "A"}
	var b Car = Car{2017, "B"}
	var i Vehicle = &a
	i.print()
	i.setName("Aa")
	i.print()

	c := i.(*Car)
	c.print()

	i = &b
	i.print()
	i.setYear(2000)
	i.print()

}
