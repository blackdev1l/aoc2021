package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/blackdev1l/aoc2021/utils"
)

type Bingo struct {
	winningNumbers   []int
	boards           []Board
	winnerBoard      Board
	winnerNumber     int
	lastWinnerBoard  Board
	lastWinnerNumber int
}

type Board struct {
	cells       [5][5]Cell
	rowCount    [5]int
	columnCount [5]int
	won         bool
}

type Cell struct {
	value    int
	isMarked bool
}

func parseInput(lines []string) Bingo {
	bingo := Bingo{}

	winning := strings.Split(lines[0], ",")

	for _, num := range winning {
		value, err := strconv.Atoi(string(num))
		utils.Check(err)
		bingo.winningNumbers = append(bingo.winningNumbers, value)
	}

	for i := 1; i < len(lines); i += 6 {
		board := Board{}

		for j := 1; j < 6; j++ {
			//fmt.Println(lines[i+j])

			line := strings.Split(lines[i+j], " ")
			var count int
			for _, l := range line {

				if l == " " || l == "" {
					continue
				}
				val, err := strconv.Atoi(l)
				//fmt.Println(val)
				utils.Check(err)
				board.cells[j-1][count] = Cell{val, false}
				//fmt.Println(board)
				count++

			}
			count = 0
		}
		bingo.boards = append(bingo.boards, board)
	}

	return bingo
}

func (bingo *Bingo) SetMarked(number int) {
	boards := bingo.boards
	for i := 0; i < len(boards); i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 5; k++ {

				if boards[i].cells[j][k].value == number {
					boards[i].cells[j][k].isMarked = true
					boards[i].rowCount[k]++
					boards[i].columnCount[j]++
					bingo.checkIfBoardIsWinner(&boards[i], number)
				}
			}
		}
		bingo.boards = boards

	}
}

func (bingo *Bingo) ThrowNumber() {
	for _, number := range bingo.winningNumbers {
		bingo.SetMarked(number)

	}
}

func (bingo *Bingo) checkIfBoardIsWinner(board *Board, number int) {
	if board.won {
		return
	}

	win := false

	for i := 0; i < 5; i++ {
		if board.columnCount[i] == 5 || board.rowCount[i] == 5 {
			win = true
		}
	}

	if win {
		fmt.Println("WINNER BOARD!")
		printBoard(*board)
		board.won = true
		if bingo.winnerNumber == 0 {
			bingo.winnerBoard = *board
			bingo.winnerNumber = number
		}

		bingo.lastWinnerBoard = *board
		bingo.lastWinnerNumber = number
		return
	}

}

func printBoard(board Board) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board.cells[i][j].isMarked {
				fmt.Printf("(%v) ", board.cells[i][j].value)
			} else {
				fmt.Printf("[%v] ", board.cells[i][j].value)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (bingo *Bingo) GetResultFromBoard() int {
	count := 0
	board := bingo.winnerBoard
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board.cells[i][j].isMarked {
				count += board.cells[i][j].value

			}
		}
	}

	fmt.Printf("count is %v winnderNumber is %v", count, bingo.winnerNumber)
	return count * bingo.winnerNumber
}

func (bingo *Bingo) GetResultFromLastBoard() int {
	count := 0
	board := bingo.lastWinnerBoard
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !board.cells[i][j].isMarked {
				count += board.cells[i][j].value

			}
		}
	}

	fmt.Printf("count is %v winnderNumber is %v", count, bingo.winnerNumber)
	return count * bingo.lastWinnerNumber
}

func main() {
	lines := utils.ReadFIleAsStringLines("input.txt")
	bingo := parseInput(lines)
	bingo.ThrowNumber()
	res := bingo.GetResultFromLastBoard()
	fmt.Printf("result is %v\n", res)
}
