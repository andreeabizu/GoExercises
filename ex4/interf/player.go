package interf

import "fmt"

type Player interface {
	play() string
	stop() string
}

type Person interface {
	eat() string
	walk() string
}

type Both interface {
	Player
	Person
}

type Man struct {
	name string
}

func (m Man) eat() string {
	s := "Man " + m.name + " eats-"
	return s
}

func (m Man) walk() string {
	s := "Man " + m.name + " walks-"
	return s
}

func (m Man) play() string {
	s := "Man " + m.name + " plays-"
	return s
}

func (m Man) stop() string {
	s := "Man " + m.name + " stops-"
	return s
}

type Woman struct {
	name string
}

func (w Woman) eat() string {
	s := "Woman " + w.name + " eats-"
	return s
}

func (w Woman) walk() string {
	s := "Woman " + w.name + " walks-"
	return s
}

func (w Woman) play() string {
	s := "Woman " + w.name + " plays-"
	return s
}

func (w Woman) stop() string {
	s := "Woman " + w.name + " stops-"
	return s
}

func Initialize() {
	m := make([]Both, 5)

	m[0] = Man{"Paul"}
	m[1] = Woman{"Andreea"}
	m[2] = Woman{"Maria"}
	m[3] = Man{"Mihai"}
	m[4] = Man{"Ion"}

	for _, val := range m {

		fmt.Println(val.eat(), val.play(), val.stop(), val.walk())
	}

	switch v := m[0].(type) {
	case Man:
		fmt.Println(v.name + " Man")
	case Woman:
		fmt.Println(v.name + " Woman")
	default:
		fmt.Println("--")
	}
}
