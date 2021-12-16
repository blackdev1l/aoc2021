package main

import (
	"testing"

	"github.com/blackdev1l/aoc2021/utils"
)

func TestCheckSyntax(t *testing.T) {
	line := "{([(<{}[<>[]}>{[]{[(<()>"
	res := checkSyntax(line)
	if res != 1197 {
		t.Errorf("res is %v but shoul be 1197\n", res)
	}
}

func TestExample(t *testing.T) {
	lines := utils.ReadFIleAsStringLines("example.txt")
	var res = 0
	for _, line := range lines {
		res += checkSyntax(line)
	}
	if res != 26397 {
		t.Errorf("res is %v but shoul be 1197\n", res)
	}
}
