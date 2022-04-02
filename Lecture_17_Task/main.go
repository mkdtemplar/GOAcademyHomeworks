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

func CompareCards(faceOfCard1, suitOfCard1, faceOfCard2, suitOfCard2 string, deck []DeckOfCards) DeckOfCards{

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
		fmt.Println("Bigger card is: ")
		return DeckOfCards{deck: Card{
			face: deck[faceIndex1].deck.face,
			suit: deck[suitIndex1].deck.suit,
		}}

	} else if faceIndex1 < faceIndex2 {
		fmt.Println("Bigger card is: ")
		return DeckOfCards{deck: Card{
			face: deck[faceIndex2].deck.face,
			suit: deck[suitIndex2].deck.suit,
		}}

	} else if faceIndex1 == faceIndex2 && suitIndex1 > suitIndex2 {
		fmt.Println("Bigger card is: ")
		return DeckOfCards{deck: Card{
			face: deck[faceIndex1].deck.face,
			suit: deck[suitIndex2].deck.suit,
		}}

	} else if faceIndex1 == faceIndex2 && suitIndex1 < suitIndex2 {
		fmt.Println("Bigger card is: ")
		return DeckOfCards{deck: Card{
			face: deck[faceIndex2].deck.face,
			suit: deck[suitIndex1].deck.suit,
		}}
	}

	if faceIndex1 == faceIndex2 && suitIndex1 == suitIndex2 {
		fmt.Println("Cards are equal")
	}
	return DeckOfCards{deck: Card{
		face: deck[faceIndex1].deck.face,
		suit: deck[suitIndex1].deck.suit,
	}}
}

func main() {

	var faceOfCard1 string
	var faceOfCard2 string
	var suitOfCard1 string
	var suitOfCard2 string
	d := make([]DeckOfCards, 52)
	
	for i := range d {
		d[i].deck = Card{face: faces[i%13], suit: suits[i/13]}
		fmt.Print(d[i])
		if i < len(d)-1 {
			fmt.Print(",")
			if i%4 == 0 && i > 3 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
	fmt.Print("Please enter face of the card ", 1, " (Deuce, Three, Four, Five, Six, Seven, Eight, Nine, Ten,Jack, Queen, King, Ace) : ")
	fmt.Scan(&faceOfCard1)

	fmt.Print("Please enter suit of the card ", 1, " (Clubs, Diamonds, Hearts, Spades) : ")
	fmt.Scan(&suitOfCard1)

	fmt.Print("Please enter face of the card ", 2, " (Deuce, Three, Four, Five, Six, Seven, Eight, Nine, Ten,Jack, Queen, King, Ace) : ")
	fmt.Scan(&faceOfCard2)

	fmt.Print("Please enter suit of the card ", 2, " (Clubs, Diamonds, Hearts, Spades) : ")
	fmt.Scan(&suitOfCard2)
	
	fmt.Println()
	
	fmt.Println(CompareCards(faceOfCard1, suitOfCard1, faceOfCard2, suitOfCard2, d))
	
	fmt.Println()

}
