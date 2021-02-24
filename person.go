package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

func print(p Person) {
	fmt.Println(p.Name, p.Age, p.Address)
}

func getDetails(p Person) string {
	return p.Name + " " + p.Address
}

func isOver20(p Person) (bool, error) {
	if p.Age <= 20 {
		return false, errors.New("This person is under 20 years old ")
	}

	return true, nil
}

func main() {
	var list = make([]Person, 3)
	p1 := Person{"Andreea", 22, "Timisoara"}
	p2 := Person{Name: "Elena", Age: 18}
	p3 := Person{Name: "Paul", Age: 30}
	list[0] = p1
	list[1] = p2
	list[2] = p3

	for _, v := range list {
		print(v)
		fmt.Println(getDetails(v))
		_, err := isOver20(v)
		if err != nil {
			fmt.Println(err)
		}
	}
}
