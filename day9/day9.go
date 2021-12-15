package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/blackdev1l/aoc2021/utils"
)

type Point struct {
	x, y int
}

func (point *Point) key() string {
	return fmt.Sprintf("%v%v", point.y, point.x)
}

func main() {
	input := readInput(utils.ReadFIleAsStringLines("input.txt"))
	points := []Point{}

	lowestValues := findLowest(input, &points)
	res := sumLowest(lowestValues)

	fmt.Printf("result part 1 is %v\n", res)

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
	fmt.Printf("result part 2 is %v\n", result)

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

func findBasin(input [][]int, point Point, basin *map[string]Point) *map[string]Point {
	_, isPresent := (*basin)[point.key()]
	if isPresent {
		return basin
	}

	(*basin)[point.key()] = point

	x := point.x
	y := point.y

	var adj = 0
	if y != 0 {
		adj = input[y-1][x]
		if adj != 9 {
			new := Point{x: x, y: y - 1}
			(*basin)[point.key()] = new
			findBasin(input, new, basin)
		}
	}

	if y+1 < len(input) {
		adj = input[y+1][x]
		if adj != 9 {
			new := Point{x: x, y: y + 1}
			(*basin)[point.key()] = new
			findBasin(input, new, basin)
		}

	}

	if x != 0 {
		adj = input[y][x-1]
		if adj != 9 {
			new := Point{x: x - 1, y: y}
			(*basin)[point.key()] = new
			findBasin(input, new, basin)
		}
	}

	if x+1 < len(input[0]) {
		adj = input[y][x+1]
		if adj != 9 {
			new := Point{x: x + 1, y: y}
			(*basin)[point.key()] = new
			findBasin(input, new, basin)
		}
	}

	return basin
}
