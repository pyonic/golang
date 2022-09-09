package main

import "fmt"

type bird interface {
	fly()
	swim()
}

type eagle struct {
	name     string
	location string
}

func (e eagle) fly() {
	fmt.Println(e.name + " is flying!")
}

func (e eagle) swim() {
	fmt.Println(e.name + " is swimming!")
}

func main() {
	var flappy bird = eagle{"Flappy", "Canada"}
	flappy.fly()
	flappy.swim()
}
