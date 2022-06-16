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
	var i, j int
	fmt.Scan(&i, &j)
	for !g.Board.ValidateTheMove(i, j) {
		fmt.Println("Please enter a vald move")
		fmt.Scan(&i, &j)
	}
	g.Board.ApplyTheMove(i, j, g.PlayerController.Curr.Symbol)
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
		fmt.Println(g.PlayerController.Curr.Name, "  Wins")
	} else {
		fmt.Println("Game drawn")
	}
}
func (g Game) RunGame() {
	for g.State == Ongoing {
		g.TakeAValidMove()
		g.PlayerController.Curr = g.PlayerController.IteratePlayer()
		g.Board.PrintBoard()
		g.State = g.DetermineState()
	}
	g.DisplayResults()
}
