package main

import "fmt"

func main() {
	cards := newDeck()  // a new deck of cards created
	fmt.Println(cards.toString())  // convert new deck of cards into a comma separated string
}
