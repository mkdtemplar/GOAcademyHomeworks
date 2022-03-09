package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	face  string
	suit  string
	cards *Card // next
}

type DeckOfCards struct {
	deck *Card // head
}

var faces = [13]string{"Deuce", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack",
	"Queen", "King", "Ace"}
var suits = [4]string{"Clubs", "Diamonds", "Hearts", "Spades"}

func InitializeDeck(dc *[52]DeckOfCards) {

	fmt.Println("Printing deck of cards in asccedenting order:")
	for i := range dc {
		dc[i].deck = &Card{face: faces[i%13], suit: suits[i/13]}
		fmt.Print(dc[i].deck.face, "-", dc[i].deck.suit, ", ")

		if i%4 == 0 && i > 0 {
			fmt.Println()

		}
	}
	fmt.Println()
	fmt.Println()
}

func shuffle(dc *[52]DeckOfCards) {

	fmt.Println("Printing cards in random order shuffled: ")
	fmt.Println()
	for i := range dc {

		second := rand.Intn(52)
		temp := dc[i]
		dc[i] = dc[second]
		dc[second] = temp
		fmt.Print(dc[i].deck.face, "-", dc[i].deck.suit, ", ")

		if i%4 == 0 {
			fmt.Println()
		}
	}

}

func deal(dealCard *[52]DeckOfCards) error {

	if dealCard == nil {
		return fmt.Errorf("list is empty")
	}

	dealCard[0].deck = dealCard[0].deck.cards

	return nil
}

func  DealOneCard(card *[52]DeckOfCards)  {

	crd := card[0:]
	for len(crd) != 0{
		fmt.Println(crd[0].deck.face, "-", crd[0].deck.suit, " ")
			crd = append(crd[:0], crd[1:]...)
	}
}

func main() {

	cardDeck := &[52]DeckOfCards{}
	InitializeDeck(cardDeck)
	fmt.Println()
	shuffle(cardDeck)

	fmt.Println()
	DealOneCard(cardDeck)
	//deal(cardDeck)
}
