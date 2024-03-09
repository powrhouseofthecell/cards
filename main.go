package main

func main() {
	cards := new_deck()
	cards.shuffle()
	cards.print()
	// strings := cards.to_string()
	// fmt.Println(strings)
}
