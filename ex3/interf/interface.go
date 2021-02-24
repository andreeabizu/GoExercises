package interf

import "fmt"

type Animal interface {
	Eat()
	Walk()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (d Dog) Eat() {
	fmt.Println(d.Name, "eats")
}

func (d Dog) Walk() {
	fmt.Println(d.Name, "walks")
}

func (c Cat) Eat() {
	fmt.Println(c.Name, "eats")
}

func (c Cat) Walk() {
	fmt.Println(c.Name, "walks")
}
