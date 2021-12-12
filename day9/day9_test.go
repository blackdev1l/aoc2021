package main

import (
	"fmt"
	"testing"

	"github.com/blackdev1l/aoc2021/utils"
)

func TestParseInput(t *testing.T) {
	file := utils.ReadFIleAsStringLines("example.txt")
	res := readInput(file)
	fmt.Println(res)
	if res[0][0] != 2 {
		t.Errorf("error input is %v ", res[0][0])
	}
}

func TestExample(t *testing.T) {
	input := readInput(utils.ReadFIleAsStringLines("example.txt"))
	points := []Point{}

	lowestValues := findLowest(input, &points)
	res := sumLowest(lowestValues)

	fmt.Printf("result is %v\n", res)
}
