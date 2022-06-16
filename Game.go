package main

import "fmt"

const Won = "won"
const Ongoing = "ongoing"
const Drawn = "drawn"

type IGame interface {
	RunGame()
	TakeAValidMove()
	DisplayResults()
}

type Game struct {
	Board
	State string
	PlayerController
}

func (g *Game) TakeAValidMove() {
	var move Move
	move = g.PlayerController.Curr.makeamove()
	for !g.Board.ValidateTheMove(move.I, move.J) {
		fmt.Println("Please enter a valid move")
		move = g.PlayerController.Curr.makeamove()
	}
	g.Board.ApplyTheMove(move.I, move.J, g.PlayerController.Curr.getSymbol())
}

//return enum or constant
func (g Game) DetermineState() string {
	if g.Board.board[0][0] == 'X' || g.Board.board[0][0] == 'O' {
		if g.Board.board[0][0] == g.Board.board[0][1] && g.Board.board[0][0] == g.Board.board[0][2] {
			return Won
		}
		if g.Board.board[0][0] == g.Board.board[1][0] && g.Board.board[0][0] == g.Board.board[2][0] {
			return Won
		}
	}
	if g.Board.board[1][1] == 'X' || g.Board.board[1][1] == 'O' {
		if g.Board.board[1][1] == g.Board.board[0][1] && g.Board.board[1][1] == g.Board.board[2][1] {
			return Won
		}
		if g.Board.board[1][1] == g.Board.board[1][0] && g.Board.board[1][1] == g.Board.board[1][2] {
			return Won
		}
		if g.Board.board[1][1] == g.Board.board[0][0] && g.Board.board[1][1] == g.Board.board[2][2] {
			return Won
		}
		if g.Board.board[1][1] == g.Board.board[2][0] && g.Board.board[1][1] == g.Board.board[0][2] {
			return Won
		}
	}
	if g.Board.board[2][2] == 'X' || g.Board.board[2][2] == 'O' {
		if g.Board.board[2][2] == g.Board.board[2][0] && g.Board.board[2][2] == g.Board.board[2][1] {
			return Won
		}
		if g.Board.board[2][2] == g.Board.board[0][2] && g.Board.board[2][2] == g.Board.board[1][2] {
			return Won
		}
	}
	if g.Board.IsFilled() {
		return Drawn
	}
	return Ongoing
}

//use enum/const for won
func (g Game) DisplayResults() {
	if g.State == Won {
		fmt.Println(g.PlayerController.Curr.getName(), "  Wins")
	} else {
		fmt.Println("Game drawn")
	}
}
func (g Game) RunGame() {
	for g.State == Ongoing {
		g.PlayerController.Curr = g.PlayerController.IteratePlayer()
		g.TakeAValidMove()
		g.Board.PrintBoard()
		g.State = g.DetermineState()
	}
	g.DisplayResults()
}
