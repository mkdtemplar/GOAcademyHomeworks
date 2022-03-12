package main

import (
	cardDraw "Lecture9/CardDraw"
	"Lecture9/CardGame"
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
