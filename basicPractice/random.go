package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(i, j int) (z int) {

	z = rand.Intn(j-i) + i

	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x := random(1, 7)

	fmt.Println(x)
}
