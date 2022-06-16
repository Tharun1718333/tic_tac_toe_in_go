package main

import "fmt"

const won = "won"
const ongoing = "ongoing"
const drawn = "drawn"

type GameFunctions interface {
	RunGame()
	TakeAValidMove()
	DisplayResults()
}

type Game struct {
	Board
	State string
	CurrentPlayer
}

func (g *Game) TakeAValidMove() {
	var i, j int
	fmt.Scan(&i, &j)
	for !g.Board.ValidateTheMove(i, j) {
		fmt.Println("Please enter a valid move")
		fmt.Scanf("%d %d", &i, &j)
	}
	g.Board.ApplyTheMove(i, j, g.Curr.Symbol)
}

// return enum or constant
func (g Game) DetermineState() string {
	if g.Board.GameBoard[0][0] == 'X' || g.Board.GameBoard[0][0] == 'O' {
		if g.GameBoard[0][0] == g.GameBoard[0][1] && g.GameBoard[0][0] == g.GameBoard[0][2] {
			return won
		}
		if g.GameBoard[0][0] == g.GameBoard[1][0] && g.GameBoard[0][0] == g.GameBoard[2][0] {
			return won
		}
	}
	if g.Board.GameBoard[1][1] == 'X' || g.Board.GameBoard[1][1] == 'O' {
		if g.GameBoard[1][1] == g.GameBoard[0][1] && g.GameBoard[1][1] == g.GameBoard[2][1] {
			return won
		}
		if g.GameBoard[1][1] == g.GameBoard[1][0] && g.GameBoard[1][1] == g.GameBoard[1][2] {
			return won
		}
		if g.GameBoard[1][1] == g.GameBoard[0][0] && g.GameBoard[1][1] == g.GameBoard[2][2] {
			return won
		}
		if g.GameBoard[1][1] == g.GameBoard[2][0] && g.GameBoard[1][1] == g.GameBoard[0][2] {
			return won
		}
	}
	if g.Board.GameBoard[2][2] == 'X' || g.Board.GameBoard[2][2] == 'O' {
		if g.GameBoard[2][2] == g.GameBoard[2][0] && g.GameBoard[2][2] == g.GameBoard[2][1] {
			return won
		}
		if g.GameBoard[2][2] == g.GameBoard[0][2] && g.GameBoard[2][2] == g.GameBoard[1][2] {
			return won
		}
	}
	if g.Board.IsFilled() {
		return drawn
	}
	return ongoing
}

// use enum/const for won
func (g Game) DisplayResults() {
	if g.State == won {
		fmt.Println(g.CurrentPlayer.Curr.Name, "  Wins")
	} else {
		fmt.Println("Game drawn")
	}
}
func (g Game) RunGame() {
	for g.State == ongoing {
		g.TakeAValidMove()
		g.CurrentPlayer.Curr = g.IteratePlayer()
		g.Board.PrintBoard()
		g.State = g.DetermineState()
	}
	g.DisplayResults()
}
