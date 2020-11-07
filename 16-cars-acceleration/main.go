package main

import (
	"fmt"
)

type vehicle interface {
	accelerate()
}

// Car
type car struct {
	model string
	color string
}

func (c car) accelerate() {
	fmt.Println("Accelrating?")
}

// Toyota
type toyota struct {
	model string
	color string
	speed int
}

func (t toyota) accelerate() {
	fmt.Println("I am toyota, I accelerate fast?")
}

func doSomething(v vehicle) {
	fmt.Println(v)

}
func main() {
	c1 := car{"suzuki", "blue"}
	t1 := toyota{"Toyota", "Red", 100}
	c1.accelerate()
	t1.accelerate()
	doSomething(c1)
	doSomething(t1)
}
