package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	fmt.Println("TEST")
	
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}

func newCard() string{
	return "Five of Diamonds"
}
