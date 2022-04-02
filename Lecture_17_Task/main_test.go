package main

import (
	"testing"
)

type DeckCard struct {
	card1    Card
	card2    Card
	expected Card
}

var addResults = []DeckCard{
	{card1: Card{face: "Deuce", suit: "Clubs"}, card2: Card{face: "Deuce", suit: "Clubs"}, expected: Card{face: "Deuce", suit: "Clubs"}},
	{card1: Card{face: "Three", suit: "Clubs"}, card2: Card{face: "Three", suit: "Clubs"}, expected: Card{face: "Three", suit: "Clubs"}},
	{card1: Card{face: "Ace", suit: "Clubs"}, card2: Card{face: "Deuce", suit: "Clubs"}, expected: Card{face: "Ace", suit: "Clubs"}},
	{card1: Card{face: "Three", suit: "Clubs"}, card2: Card{face: "King", suit: "Clubs"}, expected: Card{face: "King", suit: "Clubs"}},
	{card1: Card{face: "Six", suit: "Clubs"}, card2: Card{face: "Deuce", suit: "Spades"}, expected: Card{face: "Six", suit: "Clubs"}},
	{card1: Card{face: "Jack", suit: "Clubs"}, card2: Card{face: "Jack", suit: "Diamonds"}, expected: Card{face: "Jack", suit: "Clubs"}},
	{card1: Card{face: "Ten", suit: "Clubs"}, card2: Card{face: "Deuce", suit: "Clubs"}, expected: Card{face: "Ten", suit: "Clubs"}},
	{card1: Card{face: "Seven", suit: "Diamonds"}, card2: Card{face: "Three", suit: "Clubs"}, expected: Card{face: "Seven", suit: "Diamonds"}},
	{card1: Card{face: "Queen", suit: "Clubs"}, card2: Card{face: "Queen", suit: "Spades"}, expected: Card{face: "Queen", suit: "Clubs"}},
	{card1: Card{face: "King", suit: "Clubs"}, card2: Card{face: "Ace", suit: "Diamonds"}, expected: Card{face: "Ace", suit: "Diamonds"}},
}

func TestInput(t *testing.T) {
	for _, test := range addResults {
		result := CompareCards(test.card1.face, test.card1.suit, test.card2.face, test.card2.suit)
		if result != test.expected {
			t.Fatal("Result not ok")
		}
	}
}
