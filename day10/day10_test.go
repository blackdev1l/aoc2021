package main

import (
	"testing"

	"github.com/blackdev1l/aoc2021/utils"
)

func TestCheckSyntax(t *testing.T) {
	line := "{([(<{}[<>[]}>{[]{[(<()>"
	res, _ := checkSyntax(line)
	if res != 1197 {
		t.Errorf("res is %v but shoul be 1197\n", res)
	}
}

func TestExample(t *testing.T) {
	lines := utils.ReadFIleAsStringLines("example.txt")
	var res = 0
	completedScores := []int{}
	for _, line := range lines {
		val, stack := checkSyntax(line)
		res += val
		if val == 0 {
			completedScores = append(completedScores, completeSyntax(stack))
		}
	}
	if res != 26397 {
		t.Errorf("res is %v but shoul be 1197\n", res)
	}
	if getMiddleResult(completedScores) != 288957 {
		t.Errorf("res is %v but shoul be 288957\n", getMiddleResult(completedScores))
	}
}

func TestCompleteSyntax(t *testing.T) {
	var input = "<{([{{}}[<[[[<>{}]]]>[]]"
	_, stack := checkSyntax(input)
	res := completeSyntax(stack)
	if res != 294 {
		t.Errorf("res is %v but shoul be 294\n", res)
	}
}
