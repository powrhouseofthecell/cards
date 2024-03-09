package main

import (
	"fmt"
	"math/rand"
	"strings"
)

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

func deal(d deck, hand_size int) (deck, deck) {
	return d[:hand_size], d[hand_size:]
}

func (d deck) to_string() string {
	return strings.Join([]string(d), ",")
}

func (d deck) shuffle() {
	for i := range d {
		new_position := rand.Intn(len(d) - 1)
		d[i], d[new_position] = d[new_position], d[i]
	}
}
