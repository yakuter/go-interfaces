// https://golangbot.com/interfaces-part-2/
package main

import "fmt"

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}
	d2 = &a

	/* compilation error if the following line is
	   uncommented
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)
	*/
	//d2 = a

	// d2 = &a //This works since Describer interface
	//is implemented by Address pointer in line 22
	d2.Describe()

}
