package main

import "fmt"

// Creating own type def
type deck []string

func new_deck() deck {
	cards := []string{}

	card_suit := []string{"Spades", "Dimonds", "Clubs", "Heart"}
	card_value := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range card_suit {
		for _, value := range card_value {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver function
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
