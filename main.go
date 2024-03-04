package main

func main() {
	cards := new_deck()
	hand, remaining_hand := deal(cards, 3)
	hand.print()
	remaining_hand.print()

}

func new_card() string {
	return "Ace of Speads"
}
