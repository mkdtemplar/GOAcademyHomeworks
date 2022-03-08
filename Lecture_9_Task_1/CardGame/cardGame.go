package cardGame

import "fmt"

type Card struct {
	Face  string
	Suit  string
	cards *Card // next
}

type DeckOfCards struct {
	Deck *Card // head
}

var Faces = [13]string{"Deuce", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack",
	"Queen", "King", "Ace"}
var Suits = [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}

func InitializeDeck(dc []DeckOfCards) {

	fmt.Println("Printing deck of cards in ascending order:")
	for i := range dc {
		dc[i].Deck = &Card{Face: Faces[i%13], Suit: Suits[i/13]}
		fmt.Print(dc[i].Deck.Face, "-", dc[i].Deck.Suit, ", ")

		if i%4 == 0 && i > 0 {
			fmt.Println()

		}
	}
	fmt.Println()
	fmt.Println()
}
