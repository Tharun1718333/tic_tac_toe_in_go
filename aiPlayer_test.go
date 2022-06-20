package main

import "testing"

func TestAiPlayerMakeAMove(t *testing.T) {
	p1 := Player{"p1", 'X'}
	p2 := AiPlayer{p1}
	var board [3][3]byte
	board[0][0] = 'X'
	board[0][1] = 'X'
	gameBoard := Board{board}
	move := p2.makeamove(gameBoard)
	if move.I != 0 || move.J != 2 {
		t.Errorf("failed to make a confirmed move")
	}
	board[0][1] = 0
	gameBoard = Board{board}
	move = p2.makeamove(gameBoard)
	if move.I == 2 || move.J == 2 {
		t.Errorf("failed to make a optimal move")
	}

}
