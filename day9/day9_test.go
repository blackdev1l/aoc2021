package main

import (
	"fmt"
	"sort"
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

func TestFindBasin(t *testing.T) {
	input := readInput(utils.ReadFIleAsStringLines("example.txt"))
	point := Point{x: 9, y: 0}
	basin := make(map[string]Point)
	findBasin(input, point, &basin)
	fmt.Println(len(basin))
}

func TestPart2(t *testing.T) {
	input := readInput(utils.ReadFIleAsStringLines("example.txt"))
	points := []Point{}
	findLowest(input, &points)

	basinList := []int{}
	for _, point := range points {
		basin := make(map[string]Point)
		findBasin(input, point, &basin)
		basinList = append(basinList, len(basin))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinList)))
	result := 1
	for _, v := range basinList[:3] {
		result *= v
	}

	if result != 1134 {
		t.Errorf("wrogn result, result is %v\n", result)
	}

}
