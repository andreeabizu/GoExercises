package main

import (
	"poo/interf"
)

func main() {

	m := make(map[int]interf.Animal)
	m[0] = interf.Dog{"Bruno"}
	m[1] = interf.Dog{"Brutus"}
	m[2] = interf.Cat{"Jake"}
	m[3] = interf.Dog{"Rocky"}
	m[4] = interf.Cat{"Leo"}

	for _, v := range m {
		v.Eat()
		v.Walk()
	}
}
