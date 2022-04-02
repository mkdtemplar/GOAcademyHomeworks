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
}

func TestInput(t *testing.T) {
	for _, test := range addResults {
		result := CompareCards(test.card1.face, test.card1.suit, test.card2.face, test.card2.suit)
		if result != test.expected {
			t.Fatal("Result not ok")
		}
	}
}
