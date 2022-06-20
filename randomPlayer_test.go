package main

import "testing"

func TestRandomPlayerMakeAMove(t *testing.T) {
	p1 := Player{"p1", 'X'}
	p2 := RandomPlayer{p1}
	var gameBoard [3][3]byte
	gameBoard[0][0] = 'X'
	gameBoard[0][1] = 'X'
	gameBoard[0][2] = 'X'
	gameBoard[1][0] = 'X'
	gameBoard[1][1] = 'X'
	gameBoard[1][2] = 'X'
	gameBoard[2][0] = 'X'
	gameBoard[2][1] = 'X'
	//gameBoard[2][2] = 'X'
	board := Board{gameBoard}
	move := p2.makeamove(board)
	if move.I != 2 || move.J != 2 {
		t.Errorf("failed to get the right move")
	}

}
