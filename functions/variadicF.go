package main

import "fmt"

type Student struct {
	name string
	age  int
}

func print(students ...Student) {

	for _, val := range students {
		fmt.Println(val.name, val.age)
	}
}

func main() {

	var a, b, c, d = Student{"Mihai", 22}, Student{"Laura", 21}, Student{"Matei", 24}, Student{"Ion", 20}
	print(a, b)
	print(a, b, c)
	print(a, b, c, d)

}
