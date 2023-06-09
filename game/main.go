package game

import (
	"fmt"
	"time"
	"tutorials/tools"
)

func Domain() {
	// Make players
	dog := NewDog("Axlie")
	cat := NewCat("Binky")
	owl := NewOwl("Hanna")

	// Make the game
	game := NewGame(50)

	// Join player to the game
	game.Join(cat).Join(dog).Join(owl)

	// Loop til there is a winner
	var winner Player
	for winner == nil {
		// Move the players
		game.MovePlayers()

		// Check winner
		winner = game.CheckWinner()

		// Print the Game
		tools.ClearTerminal()
		game.Print()
		time.Sleep(time.Millisecond * 250)
	}

	// Print the winner
	fmt.Printf("%s won the game!", winner)
}
