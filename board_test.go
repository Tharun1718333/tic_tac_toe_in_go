package main

import "testing"

func TestValidateTheMove(t *testing.T) {
	var gameBoard [3][3]byte
	gameBoard[1][1] = 'X'
	board := Board{gameBoard}
	if board.ValidateTheMove(3, 3) {
		t.Errorf("validation falied for 3 3")
	}
	if board.ValidateTheMove(1, 1) {
		t.Errorf("validated even when filled")
	}
}
func TestIsFilled(t *testing.T) {
	var gameBoard [3][3]byte
	gameBoard[0][1] = 'X'
	gameBoard[1][1] = 'X'
	gameBoard[2][1] = 'X'
	gameBoard[0][2] = 'X'
	gameBoard[1][2] = 'X'
	gameBoard[2][2] = 'X'
	gameBoard[0][0] = 'X'
	gameBoard[1][0] = 'X'
	board := Board{gameBoard}
	if board.IsFilled() {
		t.Errorf("showing as filled when it is not filled")
	}
	gameBoard[2][0] = 'X'
	board = Board{gameBoard}
	if !board.IsFilled() {
		t.Errorf("showing as not filled when it is filled")
	}
}
func TestApplyTheMove(t *testing.T) {
	var gameBoard [3][3]byte
	board := Board{gameBoard}
	board.ApplyTheMove(1, 1, 'X')
	if board.board[1][1] != 'X' {
		t.Errorf("failed to apply the move 1 1")
	}
}
