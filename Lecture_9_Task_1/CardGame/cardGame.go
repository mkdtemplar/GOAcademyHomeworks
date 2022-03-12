package cardGame

import (
	"errors"
	"fmt"
)

type Card struct {
	Face string
	Suit string
}

type DeckOfCards [52]struct {
	Deck *Card // head
}

var Faces = [13]string{"Deuce", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack",
	"Queen", "King", "Ace"}

var Suits = [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}

func (d DeckOfCards) Deal() *DeckOfCards {
	for i := 0; i < len(d); i++ {
		fmt.Print(d[i].Deck.Face, "-", d[i].Deck.Suit)
		if i < len(d)-1 {
			fmt.Print(",")
			if i%4 == 0 && i > 3 {
				fmt.Println()
			}
		}

	}

	return &d
}

func (d DeckOfCards) DealOneCard() *DeckOfCards {
	sl := d[0:]

	if !d.Done() {
		fmt.Println(len(sl))
		for len(sl) != 0 {
			fmt.Println(sl[0].Deck.Face, " ", sl[0].Deck.Suit)
			sl = append(sl[:0], sl[1:]...)
		}
		fmt.Println(len(sl))
	} else {
		errors.New("Deck is empty")
	}

	return nil
}

func (d DeckOfCards) Done() bool {
	if len(d) == 0 {
		return true
	} else {
		return false
	}
}
