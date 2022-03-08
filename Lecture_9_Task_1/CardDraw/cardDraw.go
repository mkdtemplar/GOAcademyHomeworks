package cardDraw

import (
	cardGame "Lecture_9_Task_1/CardGame"
	"errors"
	"fmt"
	"math/rand"
)

type Dealer interface {
	Deal() []cardGame.DeckOfCards
	Done() bool
}

func Shuffle(dc []cardGame.DeckOfCards) {

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

type Cards []cardGame.DeckOfCards

func (crd Cards) Done() bool {

	if len(crd.Deal()) == 0 {
		return true
	} else {
		return false
	}

}

func (card Cards) Deal() []cardGame.DeckOfCards {

	card = make([]cardGame.DeckOfCards, 52)

	for i := range card {

		card[i].Deck = cardGame.Card{Face: cardGame.Faces[i%13], Suit: cardGame.Suits[i/13]}

	}
	return card
}

func DealOneCard(c Cards) error {

	c = make([]cardGame.DeckOfCards, 52)
	if len(c) == 0 {

		return errors.New("empty")
	} else {
		for len(c) != 0 {
			fmt.Println(c[0])
			c = append(c[:0], c[1:]...)
		}
	}
	return nil
}

func DrawAllCards(dealer Dealer) error {
	if len(dealer.Deal()) == 0 {
		return errors.New("card deck is empty")
	}
	for i := range dealer.Deal() {
		fmt.Print(dealer.Deal()[i])
		if i%4 == 0 {
			fmt.Println()
		}
	}
	return nil
}
