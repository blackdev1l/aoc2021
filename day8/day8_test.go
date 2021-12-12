package main

import (
	"fmt"
	"testing"

	"github.com/blackdev1l/aoc2021/utils"
)

func TestParseInput(t *testing.T) {
	input := utils.ReadFIleAsStringLines("example.txt")

	parsed := parseInput(input)
	for _, p := range parsed {
		fmt.Println(p.input)
	}
}

func TestCountOutput(t *testing.T) {
	input := utils.ReadFIleAsStringLines("example.txt")

	parsed := parseInput(input)
	result := countOutput(parsed[0])
	if result != 2 {
		t.Errorf("res is %v", result)
	}
}

func TestPart1(t *testing.T) {
	input := utils.ReadFIleAsStringLines("example.txt")

	parsed := parseInput(input)
	result := 0
	for _, v := range parsed {
		result += countOutput(v)
	}
	if result != 26 {
		t.Errorf("res is %v", result)
	}
}
