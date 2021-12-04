package main

import (
	"testing"

	"github.com/blackdev1l/aoc2021/utils"
)

var winning = []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

func getExmaple() []string {
	return utils.ReadFIleAsStringLines("example.txt")
}

func TestParseInput(t *testing.T) {
	lines := getExmaple()
	bingo := parseInput(lines)
	if len(bingo.winningNumbers) != len(winning) {
		t.Error("wrong input")
	}

}

func TestThrowNumber(t *testing.T) {
	lines := getExmaple()
	bingo := parseInput(lines)
	bingo.ThrowNumber()
	res := bingo.GetResultFromBoard()
	if res != 4512 {
		t.Errorf("res should be 4512 but is %v\n", res)
	}
}

func TestLastboardWinner(t *testing.T) {
	lines := getExmaple()
	bingo := parseInput(lines)
	bingo.ThrowNumber()
	res := bingo.GetResultFromLastBoard()
	if res != 1924 {
		t.Errorf("res should be 1924 but is %v\n", res)
	}
}
