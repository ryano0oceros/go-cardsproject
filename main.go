package main

import "fmt"

func main() {
	cards := newDeckFromFile("deleteme.txt")

	cards.shuffle()

	fmt.Println(cards.toString())
}
