package main

import "fmt"

func main() {
	cards := new_deck()
	strings := cards.to_string()
	fmt.Println(strings)
}
