package main

import (
	"fmt"
	"strconv"

	"github.com/blackdev1l/aoc2021/utils"
)

type Point struct {
	x, y int
}

func main() {
	input := readInput(utils.ReadFIleAsStringLines("input.txt"))
	points := []Point{}

	lowestValues := findLowest(input, &points)
	res := sumLowest(lowestValues)

	fmt.Printf("result is %v\n", res)

}

func readInput(input []string) [][]int {
	// var ax = len(input)
	var result = [][]int{}
	for _, line := range input {
		var lineArray = []int{}
		for _, v := range line {
			re, err := strconv.Atoi(string(v))
			utils.Check(err)
			lineArray = append(lineArray, re)
		}
		result = append(result, lineArray)
	}

	return result
}

func findLowest(input [][]int, points *[]Point) []int {
	result := []int{}
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {

			current := input[i][j]

			if i != 0 {
				if !isLowest(current, input[i-1][j]) { // up
					continue
				}
			}
			if i+1 < len(input) {
				if !isLowest(current, input[i+1][j]) { // down
					continue
				}
			}
			if j != 0 {
				if !isLowest(current, input[i][j-1]) { // left
					continue
				}
			}
			if j+1 < len(input[i]) {
				if !isLowest(current, input[i][j+1]) { // right
					continue
				}
			}
			fmt.Printf("found lowest: %v\n", current)
			result = append(result, current)
			point := Point{x: j, y: i}
			(*points) = append((*points), point)
		}
	}

	return result
}

func sumLowest(lowestList []int) int {
	result := 0
	for _, val := range lowestList {
		result += (val + 1)
	}
	return result
}

func isLowest(currentValue int, adjacent int) bool {
	return currentValue < adjacent
}
