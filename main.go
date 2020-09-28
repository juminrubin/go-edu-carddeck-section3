package main

func main() {
	// cards := newDeckFromFile("my")

	cards := newDeck()
	cards.shuffle()
	cards.print()

	//cards.saveToFile("my_cards")

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// fmt.Println("------------")
	// remainingDeck.print()

}
