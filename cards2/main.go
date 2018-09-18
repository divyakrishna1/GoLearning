package main

func main() {
	cards := deck{"ace of hearts", newCard()}
	cards = append(cards, "six of spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
