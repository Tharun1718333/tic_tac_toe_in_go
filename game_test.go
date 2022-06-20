package main

import "testing"

func TestDeterminestate(t *testing.T) {
	p1 := Player{"p1", 'X'}
	p2 := Player{"p2", 'O'}
	playerStack := []Iplayer{p1, p2}
	playerController := PlayerController{playerStack, p1}
	var gameBoard [3][3]byte
	gameBoard[0][0] = 'X'
	gameBoard[1][0] = 'X'
	gameBoard[2][0] = 'X'
	board := Board{gameBoard}
	game := Game{board, Ongoing, playerController}
	if game.DetermineState() != Won {
		t.Errorf("failed to determine when won")
	}
}
