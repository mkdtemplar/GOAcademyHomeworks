package main

import (
	"testing"

)

type Cards struct {
	faces string
	suits string
}

type DeckCard struct {
	cardDeck Cards
	expected Cards
}


var addResults = [52]DeckCard{
	{cardDeck: Cards{ faces: "Deuce", suits: "Clubs",
	}},
	{cardDeck: Cards{faces: "Three", suits: "Clubs",
	}},
	{cardDeck: Cards{faces: "Four", suits: "Clubs",
	}},
	{cardDeck: Cards{faces: "Five", suits: "Clubs",
	}},
	{expected: Cards{faces: "Deuce", suits: "Clubs",
	}},
	{expected: Cards{faces: "Three", suits: "Clubs",
	}},
}


func TestInput(t *testing.T) {
	for _, t := range addResults{
		result := CompareCards()
	}
}
