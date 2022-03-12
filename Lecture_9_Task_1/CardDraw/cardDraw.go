package cardDraw

import cardGame "Lecture_9_Task_1/CardGame"

type Dealer interface {
	Deal() *cardGame.DeckOfCards
	DealOneCard() *cardGame.DeckOfCards
	Done() bool
}

func DealOneCard(c Dealer) {

	c.DealOneCard()
}

func DrawAllCards(dealer Dealer) {
	dealer.Deal()
}
