package cardGame

import "fmt"

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
		fmt.Println(d[i].Deck.Face, "-", d[i].Deck.Suit)
	}

	return &d
	panic("implement me")
}

func (d DeckOfCards) DealOneCard() *DeckOfCards {
	sl := d[0:]

	fmt.Println(len(sl))
	for len(sl) != 0 {
		fmt.Println(sl[0].Deck.Face, " ", sl[0].Deck.Suit)
		sl = append(sl[:0], sl[1:]...)
	}
	fmt.Println(len(sl))
	return &d
	panic("implement me")
}

func (d DeckOfCards) Done() bool {
	if len(d) == 0 {
		return true
	} else {
		return false
	}
	panic("implement me")
}
