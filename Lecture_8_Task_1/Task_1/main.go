package main

import (
	"Lecture_8_Task_1/Task_1/carddraw"
	"Lecture_8_Task_1/Task_1/cardgame"
	"fmt"
)

func main() {

	dc := make([]cardgame.DeckOfCards, 52)
	var dc2 carddraw.Dealer
	dc2 = &carddraw.Cards{}

	cardgame.InitializeDeck(dc)
	fmt.Println()
	fmt.Println()

	carddraw.Shuffle(dc)
	fmt.Println()
	fmt.Println()

	carddraw.DrawAllCards(dc2)
	fmt.Println()
	fmt.Println()

}
