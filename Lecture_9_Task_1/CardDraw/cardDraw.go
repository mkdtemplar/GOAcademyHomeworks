package cardDraw

import (
	cardGame "Lecture_9_Task_1/CardGame"
	"errors"
	"fmt"
)

type Dealer interface {
	Deal() *[52]cardGame.DeckOfCards
	Done() bool
}

type Cards [52]cardGame.DeckOfCards

func (crd Cards) Done() bool {

	if len(crd.Deal()) == 0 {
		return true
	} else {
		return false
	}

}

func (card Cards) Deal() *[52]cardGame.DeckOfCards {

	for i := range card {

		card[i].Deck = &cardGame.Card{Face: cardGame.Faces[i%13], Suit: cardGame.Suits[i/13]}

	}
	return (*[52]cardGame.DeckOfCards)(&card)
}

func DealOneCard(c Dealer) error {

	sl := c.Deal()[0:len(c.Deal())]

	if len(sl) == 0 {

		return errors.New("empty")
	} else {
		for i := 0; i < len(sl); i++ {
			fmt.Print(sl[0].Deck.Face, "-", sl[0].Deck.Suit, " ")
			sl = append(sl[:i], sl[i+1:]...)
			i--
		}
	}
	return nil
}

func DrawAllCards(dealer Dealer) error {
	if len(dealer.Deal()) == 0 {
		return errors.New("card deck is empty")
	}
	for i := range dealer.Deal() {
		fmt.Print(dealer.Deal()[i].Deck.Face, "-", dealer.Deal()[i].Deck.Suit, " ")
		if i%4 == 0 {
			fmt.Println()
		}
	}
	return nil
}
