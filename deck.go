package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// creates a new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// loop through the deck of cards and print out the value of them
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deals a specific amount of cards per hand
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// creates comma separated string from list
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}