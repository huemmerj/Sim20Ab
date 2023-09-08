package main

import (
	"fmt"
)

const (
	numberOfPlayers = 1
	cardsPerPlayer  = 1
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

	// values := []string{Seven, Eight}
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

// Geerate all possible hands for all players
func generateHands(deck []Card, currentCardIndex int, currentPlayerIndex int, currentHandIndex int, game *Game, result *[]Game) {
	if currentPlayerIndex == len(game.Players) {
		copiedGame := *game
		*result = append(*result, copiedGame)
		for _, curGame := range *result {
			for _, player := range curGame.Players {
				fmt.Printf("Player %d\n", player.ID)
				for _, card := range player.Hand {
					fmt.Printf("%s %s\n", card.Value, card.Color)
				}
			}
		}
		fmt.Println()
		fmt.Println()
		fmt.Println()
		return
	}

	if currentHandIndex == cardsPerPlayer {
		generateHands(deck, currentCardIndex, currentPlayerIndex+1, 0, game, result)
		return
	}

	if currentCardIndex == len(deck) {
		return
	}
	generateHands(deck, currentCardIndex+1, currentPlayerIndex, currentHandIndex+1, game, result)

	game.Players[currentPlayerIndex].Hand[currentHandIndex] = deck[currentCardIndex]
	generateHands(deck, currentCardIndex+1, currentPlayerIndex, currentHandIndex, game, result)
}
func main() {
	deck := createDeck()

	players := make([]Player, numberOfPlayers)
	for i := 0; i < numberOfPlayers; i++ {
		players[i] = Player{ID: i + 1}
	}

	game := Game{
		Players: players,
		Deck:    deck,
	}

	for i := range game.Players {
		game.Players[i].Hand = make([]Card, cardsPerPlayer)
	}

	var allPossibleGames []Game
	generateHands(deck, 0, 0, 0, &game, &allPossibleGames)

	fmt.Printf("Total possible games: %d\n", len(allPossibleGames))
	for i, game := range allPossibleGames {
		fmt.Printf("Game %d\n", i+1)
		for _, player := range game.Players {
			fmt.Printf("Player %d\n", player.ID)
			for _, card := range player.Hand {
				fmt.Printf("%s %s\n", card.Value, card.Color)
			}
		}
	}
}
