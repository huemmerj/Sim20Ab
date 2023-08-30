package main

import (
	"fmt"
)

const (
	numberOfPlayers = 1
	cardsPerPlayer  = 5
	TotalColors     = 4
	TotalValues     = 8
	Herz            = "Herz"
	Eichel          = "Eichel"
	Blatt           = "Blatt"
	Schellen        = "Schellen"
	Seven           = "7"
	Eight           = "8"
	Nine            = "9"
	Ten             = "10"
	Under           = "Unter"
	Ober            = "Ober"
	King            = "König"
	As              = "As"
)

type Card struct {
	Value string
	Color string
	Worth int
}

type Player struct {
	ID   int
	Hand []Card
}

type Game struct {
	Players []Player
	Deck    []Card
}

func createDeck() []Card {
	deck := make([]Card, 0)

	colors := []string{Herz, Schellen, Eichel, Blatt}

	values := []string{Seven, Eight, Nine, Under, Ober, King, Ten, As}
	for _, color := range colors {
		worth := 0
		for _, value := range values {
			deck = append(deck, Card{Color: color, Value: value, Worth: worth})
			worth++
		}
	}

	return deck
}
func generateHands(deck []Card, currentHandIndex int, currentCardIndex int, currentHand []Card, result *[][]Card) {
	if currentHandIndex == cardsPerPlayer {
		fmt.Printf("currentHand: %v\n", currentHand)
		copiedHand := make([]Card, len(currentHand))
		copy(copiedHand, currentHand)
		*result = append(*result, copiedHand)
		return
	}

	if currentCardIndex == len(deck) {
		// All cards have been considered
		return
	}

	// Try not including the current card in the current hand
	generateHands(deck, currentHandIndex, currentCardIndex+1, currentHand, result)

	// Try including the current card in the current hand
	currentHand[currentHandIndex] = deck[currentCardIndex]
	generateHands(deck, currentHandIndex+1, currentCardIndex+1, currentHand, result)
	currentHand[currentHandIndex] = Card{} // Reset the slot

}

func main() {
	deck := createDeck()
	var allPossibleHands [][]Card
	generateHands(deck, 0, 0, make([]Card, cardsPerPlayer), &allPossibleHands)

	fmt.Printf("Total possible hands: %d\n", len(allPossibleHands))
}
