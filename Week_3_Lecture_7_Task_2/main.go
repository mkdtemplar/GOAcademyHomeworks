package main

import (
	"fmt"
)

type Card struct {
	face string
	suit string
}

type DeckOfCards struct {
	deck Card
}

var faces = [13]string{"Deuce", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack",
	"Queen", "King", "Ace"}
var suits = [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}

func compareCards(deck []DeckOfCards) {

	var faceOfCard1 string
	var suitOfCard1 string
	var faceOfCard2 string
	var suitOfCard2 string
	fmt.Print("Please enter face of the card ", 1, " (Deuce, Three, Four, Five, Six, Seven, Eight, Nine, Ten,Jack, Queen, King, Ace) : ")
	fmt.Scan(&faceOfCard1)

	fmt.Print("Please enter suit of the card ", 1, " (Clubs, Diamonds, Hearts, Spades) : ")
	fmt.Scan(&suitOfCard1)

	fmt.Print("Please enter face of the card ", 2, " (Deuce, Three, Four, Five, Six, Seven, Eight, Nine, Ten,Jack, Queen, King, Ace) : ")
	fmt.Scan(&faceOfCard2)

	fmt.Print("Please enter suit of the card ", 2, " (Clubs, Diamonds, Hearts, Spades) : ")
	fmt.Scan(&suitOfCard2)

	var faceIndex1 int
	var faceIndex2 int

	for i := 0; i < len(deck); i++ {
		if faceOfCard1 == deck[i].deck.face {
			faceIndex1 = i
		}
	}

	for i := 0; i < len(deck); i++ {
		if faceOfCard2 == deck[i].deck.face {
			faceIndex2 = i
		}
	}

	var suitIndex1 int
	var suitIndex2 int

	for i := 0; i < len(deck); i++ {

		if suitOfCard1 == deck[i].deck.suit {
			suitIndex1 = i
		}
	}

	for i := 0; i < len(deck); i++ {
		if suitOfCard2 == deck[i].deck.suit {

			suitIndex2 = i
		}
	}

	fmt.Println()

	if faceIndex1 > faceIndex2 {
		fmt.Println("Bigger card is: ", deck[faceIndex1].deck.face, " ", deck[suitIndex1].deck.suit)

	} else if faceIndex1 < faceIndex2 {
		fmt.Println("Bigger card is: ", deck[faceIndex2].deck.face, " ", deck[suitIndex2].deck.suit)

	} else if faceIndex1 == faceIndex2 && suitIndex1 > suitIndex2 {
		fmt.Println("Bigger card is: ", deck[faceIndex1].deck.face, " ", deck[suitIndex2].deck.suit)

	} else if faceIndex1 == faceIndex2 && suitIndex1 < suitIndex2 {
		fmt.Println("Bigger card is: ", deck[faceIndex2].deck.face, " ", deck[suitIndex1].deck.suit)
	}

	if faceIndex1 == faceIndex2 && suitIndex1 == suitIndex2 {
		fmt.Println("Cards are equal")
	}
}
func maxCard(cards []DeckOfCards) {

	fmt.Println("Max card from the slice: ", cards[len(cards)-1])
}
func main() {

	d := make([]DeckOfCards, 52)
	fmt.Println("Printing deck of all cards with face and suit")

	for i := range d {
		d[i].deck = Card{face: faces[i%13], suit: suits[i/13]}

		fmt.Print(d[i].deck)

		if i%4 == 0 && i > 0 {
			fmt.Println()
		}
	}

	deckSlice := d[5:25]
	fmt.Println()

	fmt.Println()
	fmt.Println("Slice from index 5 to 25")
	for j := range deckSlice {
		fmt.Print(deckSlice[j], " ")
		if j%4 == 0 && j > 0 {
			fmt.Println()
		}
	}
	fmt.Println()
	compareCards(d)
	fmt.Println()
	maxCard(deckSlice)
}
