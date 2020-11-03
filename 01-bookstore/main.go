// https://github.com/inancgumus/learngo/tree/master/interfaces/04-interfaces

package main

import "fmt"

// INTERFACE

type printer interface {
	print()
}

func main() {

	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubik's cube", price: 5}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik)
	store.print()

	// interface values are comparable
	// fmt.Println(store[0] == &minecraft)
	// fmt.Println(store[3] == rubik)
}

// BOOK
type book struct {
	title string
	price money
}

func (b book) print() {
	fmt.Printf("%-15s: %s\n", b.title, b.price.string())
}

func (b book) oku() {

}

// GAME
type game struct {
	title string
	price money
}

func (g *game) print() {
	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
}

func (g *game) discount(ratio float64) {
	g.price *= money(1 - ratio)
}

// PUZZLE
type puzzle struct {
	title string
	price money
}

func (p puzzle) print() {
	fmt.Printf("%-15s: %s\n", p.title, p.price.string())
}

// MONEY
type money float64

func (m money) string() string {
	return fmt.Sprintf("$%.2f", m)
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, it := range l {
		fmt.Printf("(%-10T) --> ", it)

		it.print()

		// it.discount(.5)
	}
}
