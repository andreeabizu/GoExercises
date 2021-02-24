package main

import "fmt"

type Person struct {
	name   string
	weight int
}

func avg(persons [4]Person) float32 {
	s := 0
	for i := 0; i < 4; i++ {
		s += persons[i].weight
	}

	var a float32
	a = float32(s) / 4
	return a
}

func main() {

	p := [4]Person{
		{name: "mihai", weight: 50},
		{name: "daniel", weight: 87},
		{name: "laura", weight: 50},
		{name: "andreea", weight: 52},
	}

	y := avg(p)
	fmt.Println(y)

}
