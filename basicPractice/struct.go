package main

import "fmt"

type house struct {
	noRooms int
	city    string
	price   float64
}

func main() {
	var a house
	a.noRooms = 5
	a.city = "New York"
	a.price = 82000.4

	fmt.Println(a)

	r := house{3,
		"Berlin",
		78000.56,
	}
	fmt.Println(r)

	t := house{}
	fmt.Println(t)

	var p *house
	p = &a

	fmt.Println(p.city)

	y := house{city: "L.A."}
	fmt.Println(y)

}
