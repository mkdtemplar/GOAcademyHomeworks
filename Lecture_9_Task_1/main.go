package main

import (
	cardDraw "Lecture_9_Task_1/CardDraw"
	cardGame "Lecture_9_Task_1/CardGame"
	"fmt"
)

func main() {

	dc := make([]cardGame.DeckOfCards, 52)
	var dc2 cardDraw.Dealer
	dc2 = &cardDraw.Cards{}

	cardGame.InitializeDeck(dc)
	fmt.Println()
	fmt.Println()

	cardDraw.Shuffle(dc)
	fmt.Println()
	fmt.Println()

	cardDraw.DrawAllCards(dc2)
	fmt.Println()
	fmt.Println()

	fmt.Println(dc2.Done())

}