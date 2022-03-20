package carddraw

import (
	"Lecture_8_Task_1/Task_1/cardgame"
	"fmt"
	"math/rand"
)

type Dealer interface {
	Deal() [52]cardgame.DeckOfCards
}

func Shuffle(dc []cardgame.DeckOfCards) {

	fmt.Println("Printing cards in random order shuffled: ")
	fmt.Println()
	for i := range dc {

		second := rand.Intn(52)
		temp := dc[i]
		dc[i] = dc[second]
		dc[second] = temp
		fmt.Print(dc[i].Deck.Face, "-", dc[i].Deck.Suit, ", ")

		if i%4 == 0 {
			fmt.Println()
		}
	}
}

type Cards [52]cardgame.DeckOfCards

func (card Cards) Deal() [52]cardgame.DeckOfCards {
	for i := range card {

		card[i].Deck = cardgame.Card{Face: cardgame.Faces[i%13], Suit: cardgame.Suits[i/13]}

	}
	return card
}

func DrawAllCards(dealer Dealer) {
	// call the dealer's Draw() method, until you reach a nil Card

	for i := range dealer.Deal() {
		fmt.Print(dealer.Deal()[i])
		if i%4 == 0 {
			fmt.Println()
		}
	}
}
