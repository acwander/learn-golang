package main

func main() {
	// newDeck().saveToFile("my_cards")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
