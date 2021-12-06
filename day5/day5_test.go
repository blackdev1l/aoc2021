package main

import (
	"fmt"
	"testing"

	"github.com/blackdev1l/aoc2021/utils"
)

func TestParseInput(t *testing.T) {
	lines := utils.ReadFIleAsStringLines("example.txt")
	res := parseInput(lines)
	fmt.Println(res)
	if res[0].source[0] != 0 {
		t.Error("bruh")
	}

}

func TestConnectPoint(t *testing.T) {
	input := Vent{}
	input.source = []int{0, 9}
	input.target = []int{5, 9}
	matrix := [1000][1000]int{}
	connectPoints(input, &matrix, false)

	input = Vent{}
	input.source = []int{2, 2}
	input.target = []int{2, 1}
	connectPoints(input, &matrix, false)
	printDiagram(matrix)
}

func TestOverlap(t *testing.T) {
	lines := utils.ReadFIleAsStringLines("example.txt")
	vents := parseInput(lines)
	matrix := [1000][1000]int{}
	for _, v := range vents {
		connectPoints(v, &matrix, false)
	}
	res := findOverlap(&matrix)
	printDiagram(matrix)
	if res != 5 {
		t.Errorf("res is %v but should be 5", res)
	}
}

func TestConnectPointPart2(t *testing.T) {
	input := Vent{}
	input.source = []int{5, 5}
	input.target = []int{8, 2}
	matrix := [1000][1000]int{}
	connectPoints(input, &matrix, true)

	// input = Vent{}
	// input.source = []int{0, 0}
	// input.target = []int{9, 9}
	//connectPoints(input, &matrix, true)
	printDiagram(matrix)
}

func TestPart2(t *testing.T) {
	lines := utils.ReadFIleAsStringLines("example.txt")
	vents := parseInput(lines)
	matrix := [1000][1000]int{}
	for _, v := range vents {
		connectPoints(v, &matrix, true)
	}
	printDiagram(matrix)
	res := findOverlap(&matrix)
	if res != 12 {
		t.Errorf("res is %v but should be 12", res)
	}
}
