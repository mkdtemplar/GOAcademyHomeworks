package cardDraw

import (
	//"errors"
	//"fmt"
	cardGame "Lecture_9_Task_1/CardGame"
)

type Dealer interface {
	Deal() *cardGame.DeckOfCards
	DealOneCard() *cardGame.DeckOfCards
	Done() bool
}

func DealOneCard(c Dealer) error {

	c.DealOneCard()
	return nil
}

func DrawAllCards(dealer Dealer) error {
	dealer.Deal()
	return nil
}
