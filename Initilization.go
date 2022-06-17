package main

func main() {
	player1 := Player{"player1", 'X'}
	player2 := Player{"player2", 'O'}
	randomPlayer1 := RandomPlayer{player1}
	optimalPlayer2 := AiPlayer{player2}
	var PlayerStack []Iplayer
	PlayerStack = append(PlayerStack, randomPlayer1)
	PlayerStack = append(PlayerStack, optimalPlayer2)
	var TicTacToeBoard [3][3]byte
	board1 := Board{TicTacToeBoard}
	PlayersIterator := PlayerController{PlayerStack, PlayerStack[1]}
	GameInstance1 := Game{board1, "ongoing", PlayersIterator}
	GameInstance1.RunGame()
}
