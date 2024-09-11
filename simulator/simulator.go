package main

import (
	"fmt"
)

const (
	numberOfPlayers = 2
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

	// colors := []string{Herz, Eichel, Blatt}
	colors := []string{Herz, Schellen, Eichel, Blatt}

	// values := []string{Seven, Eight, Nine, Ten}
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
		copiedGame := *game // Create a shallow copy of the game
		copiedPlayers := make([]Player, len(game.Players))

		for i := range game.Players {
			copiedPlayer := Player{
				ID:   game.Players[i].ID,
				Hand: make([]Card, len(game.Players[i].Hand)),
			}
			copy(copiedPlayer.Hand, game.Players[i].Hand) // Copy the player's hand
			copiedPlayers[i] = copiedPlayer
		}

		copiedGame.Players = copiedPlayers
		*result = append(*result, copiedGame)
		return
	}

	if currentHandIndex == cardsPerPlayer {
		generateHands(deck, currentCardIndex, currentPlayerIndex+1, 0, game, result)
		return
	}

	if currentCardIndex == len(deck) {
		return
	}

	// Copy the current card from the deck to the current player's hand
	newCard := Card{
		Value: deck[currentCardIndex].Value,
		Color: deck[currentCardIndex].Color,
		Worth: deck[currentCardIndex].Worth,
	}
	game.Players[currentPlayerIndex].Hand[currentHandIndex] = newCard

	// Recursively generate the next hand
	generateHands(deck, currentCardIndex+1, currentPlayerIndex, currentHandIndex+1, game, result)

	// Backtrack: Reset the current player's hand to an empty card for the next iteration
	game.Players[currentPlayerIndex].Hand[currentHandIndex] = Card{}

	// Skip the current card and continue generating hands
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
		fmt.Println()
	}
}
