package main

import (
	cardDraw "Lecture_9_Task_1/CardDraw"
	"Lecture_9_Task_1/CardGame"
	"fmt"
)

func main() {

	dc := &cardGame.DeckOfCards{}

	for i := range dc {
		dc[i].Deck = &cardGame.Card{Face: cardGame.Faces[i%13], Suit: cardGame.Suits[i/13]}
	}

	var dc2 cardDraw.Dealer = dc

	cardDraw.DrawAllCards(dc2)
	fmt.Println()

	fmt.Println(cardDraw.DealOneCard(dc2))

	fmt.Println()

}
