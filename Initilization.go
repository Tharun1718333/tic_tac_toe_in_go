package main

func main() {
	player1 := Player{"player1", 'X'}
	player2 := Player{"player2", 'O'}
	var PlayerStack []Player
	PlayerStack = append(PlayerStack, player1)
	PlayerStack = append(PlayerStack, player2)
	var TicTacToeBoard [3][3]byte
	board1 := Board{TicTacToeBoard}
	PlayersIterator := PlayerController{PlayerStack, PlayerStack[0]}
	GameInstance1 := Game{board1, "ongoing", PlayersIterator}
	GameInstance1.RunGame()
}
