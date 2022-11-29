package main

import "fmt"

func main() {
	cards := newDeckFromFile("deleteme.txt")

	fmt.Println(cards.toString())
}
