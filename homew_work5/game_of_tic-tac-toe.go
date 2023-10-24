package home_work5

import "fmt"

type GameBoard map[int]string

type TicCatToe struct{}

func (game *TicCatToe) StarGame() {
	board := make(GameBoard)
	board.initializeBoard()
	player := "X"
	winner := ""

	for !game.isGameOver(board) {
		fmt.Println("Current state of the game:")
		board.displayBoard()

		fmt.Printf("Player %s, enter row and column (eg '1 2'): ", player)
		var row, col int
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err != nil || row < 0 || row > 2 || col < 0 || col > 2 || board[row*3+col] != "" {
			fmt.Println("!!!WARNING!!!\nIncorrect input.\n Try again.")
			continue
		}

		board[row*3+col] = player
		if game.hasWinner(board, player) {
			winner = player
			break
		}

		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}

	fmt.Println("Current state of the game:")
	board.displayBoard()

	if winner != "" {
		fmt.Printf("!!! Player %s has won !!!\n", winner)
	} else {
		fmt.Println("!!! Nobody's !!!")
	}
}

func (game *TicCatToe) isGameOver(board GameBoard) bool {
	return len(board.emptyCells()) == 0 || game.hasWinner(board, "X") || game.hasWinner(board, "O")
}

func (game *TicCatToe) hasWinner(board GameBoard, player string) bool {
	winningCombinations := [8][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, combo := range winningCombinations {
		if board[combo[0]] == player && board[combo[1]] == player && board[combo[2]] == player {
			return true
		}
	}
	return false
}

func (gameBoard *GameBoard) initializeBoard() {
	for i := 0; i < 9; i++ {
		(*gameBoard)[i] = ""
	}
}

func (board GameBoard) emptyCells() []int {
	var empty []int
	for i, cell := range board {
		if cell == "" {
			empty = append(empty, i)
		}
	}
	return empty
}

func (board GameBoard) displayBoard() {
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println()
			fmt.Println("---------")
		}
		if board[i] == "" {
			fmt.Print(" . ")
		} else {
			fmt.Printf(" %s ", board[i])
		}
	}
	fmt.Println()
}
