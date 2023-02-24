package main

import (
	"fmt"
)

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("Woof! My name is %s.", d.Name)
}

func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}
func main() {
	d := Dog{Name: "Shaz"}
	MakeAnimalSpeak(d)
}
