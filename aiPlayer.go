package main

//plan
/*
are you winning() apply the winning move
is the opponent winning() block out the winning move
see if you can gain on an existing move() apply the gaining move
apply the move with a preference order()

the top two functions can be combined into 1
*/
// 158 = 2 'O' and 176 = 2 'X'
// check for 8 cases of win each with 3 cases of Immediate win
import (
	"math/rand"
)

type AiPlayer struct {
	Player
}

func (p AiPlayer) getName() string {
	return p.Name
}
func (p AiPlayer) getSymbol() byte {
	return p.Symbol
}
func (p AiPlayer) makeamove(b Board) Move {
	move := p.confirmedNextMove(b)
	if move.I != 3 {
		return move
	}
	move = p.optimalMove(b)
	if move.I != 3 {
		return move
	}
	move = p.randomMove(b)
	return move
}
func (p AiPlayer) confirmedNextMove(b Board) Move {
	//Horizantl cases
	sumOfThreeBlocks := b.board[0][0] + b.board[0][1] + b.board[0][2]
	if sumOfThreeBlocks == 158 || sumOfThreeBlocks == 176 {
		if b.board[0][0] == 0 {
			return Move{0, 0}
		}
		if b.board[0][1] == 0 {
			return Move{0, 1}
		}
		if b.board[0][2] == 0 {
			return Move{0, 2}
		}
	}
	sumOfThreeBlocks = b.board[1][0] + b.board[1][1] + b.board[1][2]
	if sumOfThreeBlocks == 158 || sumOfThreeBlocks == 176 {
		if b.board[1][0] == 0 {
			return Move{1, 0}
		}
		if b.board[1][1] == 0 {
			return Move{1, 1}
		}
		if b.board[1][2] == 0 {
			return Move{1, 2}
		}
	}
	sumOfThreeBlocks = b.board[2][0] + b.board[2][1] + b.board[2][2]
	if sumOfThreeBlocks == 158 || sumOfThreeBlocks == 176 {
		if b.board[2][0] == 0 {
			return Move{2, 0}
		}
		if b.board[2][1] == 0 {
			return Move{2, 1}
		}
		if b.board[2][2] == 0 {
			return Move{2, 2}
		}
	}
	//vertical cases
	sumOfThreeBlocks = b.board[0][0] + b.board[1][0] + b.board[2][0]
	if sumOfThreeBlocks == 158 || sumOfThreeBlocks == 176 {
		if b.board[0][0] == 0 {
			return Move{0, 0}
		}
		if b.board[1][0] == 0 {
			return Move{1, 0}
		}
		if b.board[2][0] == 0 {
			return Move{2, 0}
		}
	}
	sumOfThreeBlocks = b.board[0][1] + b.board[1][1] + b.board[2][1]
	if sumOfThreeBlocks == 158 || sumOfThreeBlocks == 176 {
		if b.board[0][1] == 0 {
			return Move{0, 1}
		}
		if b.board[1][1] == 0 {
			return Move{1, 1}
		}
		if b.board[2][1] == 0 {
			return Move{2, 1}
		}
	}
	sumOfThreeBlocks = b.board[0][2] + b.board[1][2] + b.board[2][2]
	if sumOfThreeBlocks == 158 || sumOfThreeBlocks == 176 {
		if b.board[0][2] == 0 {
			return Move{0, 2}
		}
		if b.board[1][2] == 0 {
			return Move{1, 2}
		}
		if b.board[2][2] == 0 {
			return Move{2, 2}
		}
	}
	//diagonal cases
	sumOfThreeBlocks = b.board[0][0] + b.board[1][1] + b.board[2][2]
	if sumOfThreeBlocks == 158 || sumOfThreeBlocks == 176 {
		if b.board[0][0] == 0 {
			return Move{0, 0}
		}
		if b.board[1][1] == 0 {
			return Move{1, 1}
		}
		if b.board[2][2] == 0 {
			return Move{2, 2}
		}
	}
	sumOfThreeBlocks = b.board[2][0] + b.board[1][1] + b.board[0][2]
	if sumOfThreeBlocks == 158 || sumOfThreeBlocks == 176 {
		if b.board[2][0] == 0 {
			return Move{2, 0}
		}
		if b.board[1][1] == 0 {
			return Move{1, 1}
		}
		if b.board[0][2] == 0 {
			return Move{0, 2}
		}
	}
	return Move{3, 3}
}
func (p AiPlayer) optimalMove(b Board) Move {
	sumOfThreeBlocks := b.board[0][0] + b.board[0][1] + b.board[0][2]
	if sumOfThreeBlocks == p.getSymbol() {
		if b.board[0][0] == p.getSymbol() {
			return Move{0, 1}
		}
		if b.board[0][1] == p.getSymbol() {
			return Move{0, 2}
		}
		if b.board[0][2] == p.getSymbol() {
			return Move{0, 1}
		}
	}
	sumOfThreeBlocks = b.board[1][0] + b.board[1][1] + b.board[1][2]
	if sumOfThreeBlocks == p.getSymbol() {
		if b.board[1][0] == p.getSymbol() {
			return Move{1, 1}
		}
		if b.board[1][1] == p.getSymbol() {
			return Move{1, 2}
		}
		if b.board[1][2] == p.getSymbol() {
			return Move{1, 1}
		}
	}
	sumOfThreeBlocks = b.board[2][0] + b.board[2][1] + b.board[2][2]
	if sumOfThreeBlocks == p.getSymbol() {
		if b.board[2][0] == p.getSymbol() {
			return Move{2, 1}
		}
		if b.board[2][1] == p.getSymbol() {
			return Move{2, 2}
		}
		if b.board[2][2] == p.getSymbol() {
			return Move{2, 1}
		}
	}
	sumOfThreeBlocks = b.board[0][0] + b.board[1][0] + b.board[2][0]
	if sumOfThreeBlocks == p.getSymbol() {
		if b.board[0][0] == p.getSymbol() {
			return Move{1, 0}
		}
		if b.board[1][0] == p.getSymbol() {
			return Move{2, 0}
		}
		if b.board[2][0] == p.getSymbol() {
			return Move{1, 0}
		}
	}
	sumOfThreeBlocks = b.board[0][1] + b.board[1][1] + b.board[2][1]
	if sumOfThreeBlocks == p.getSymbol() {
		if b.board[0][1] == p.getSymbol() {
			return Move{1, 1}
		}
		if b.board[1][1] == p.getSymbol() {
			return Move{2, 1}
		}
		if b.board[2][1] == p.getSymbol() {
			return Move{1, 1}
		}
	}
	sumOfThreeBlocks = b.board[0][2] + b.board[1][2] + b.board[2][2]
	if sumOfThreeBlocks == p.getSymbol() {
		if b.board[0][2] == p.getSymbol() {
			return Move{1, 2}
		}
		if b.board[1][2] == p.getSymbol() {
			return Move{2, 2}
		}
		if b.board[2][2] == p.getSymbol() {
			return Move{1, 2}
		}
	}
	sumOfThreeBlocks = b.board[0][0] + b.board[1][1] + b.board[2][2]
	if sumOfThreeBlocks == p.getSymbol() {
		if b.board[0][0] == p.getSymbol() {
			return Move{1, 1}
		}
		if b.board[1][1] == p.getSymbol() {
			return Move{2, 2}
		}
		if b.board[2][2] == p.getSymbol() {
			return Move{1, 1}
		}
	}
	sumOfThreeBlocks = b.board[2][0] + b.board[1][1] + b.board[0][2]
	if sumOfThreeBlocks == p.getSymbol() {
		if b.board[2][0] == p.getSymbol() {
			return Move{1, 1}
		}
		if b.board[1][1] == p.getSymbol() {
			return Move{0, 2}
		}
		if b.board[0][2] == p.getSymbol() {
			return Move{1, 1}
		}
	}
	return Move{3, 3}
}
func (p AiPlayer) randomMove(b Board) Move {
	var allPossiblePositions []Move
	for l := 0; l < 3; l++ {
		for m := 0; m < 3; m++ {
			if b.board[l][m] == 0 {
				allPossiblePositions = append(allPossiblePositions, Move{l, m})
			}
		}
	}
	return allPossiblePositions[rand.Intn(len(allPossiblePositions))]
}
