package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/blackdev1l/aoc2021/utils"
)

type Vent struct {
	source []int
	target []int
}

func parseInput(list []string) []Vent {
	result := []Vent{}
	for _, line := range list {
		parsedLine := strings.Split(line, "->")
		sourceString := strings.Split(parsedLine[0], ",")
		targetString := strings.Split(parsedLine[1], ",")

		source := ArrayStringToArrayInt(sourceString)
		target := ArrayStringToArrayInt(targetString)
		result = append(result, Vent{source: source, target: target})

	}
	return result
}

func ArrayStringToArrayInt(arr []string) []int {

	result := []int{}
	for _, el := range arr {
		el = strings.ReplaceAll(el, " ", "")
		i, err := strconv.Atoi(el)
		utils.Check(err)
		result = append(result, i)
	}
	return result
}

func printDiagram(matrix [1000][1000]int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("|%v", i)
		for j := 0; j < 10; j++ {
			if matrix[i][j] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%v", matrix[i][j])
			}

		}
		fmt.Printf("|")
		fmt.Println()
	}

}

func connectPoints(vent Vent, matrix *[1000][1000]int, part2 bool) {
	src := vent.source
	trg := vent.target
	// segmento verticale
	var axys = 0
	// diagonale
	if src[0] != trg[0] && src[1] != trg[1] {
		if part2 {
			connectDiagonally(vent, matrix)
		}

		return
	}

	// segmento orizzontale
	if src[0] == trg[0] {
		axys = 1
	}

	var start, end int
	if src[axys] > trg[axys] {
		start = trg[axys]
		end = src[axys]
	} else {
		start = src[axys]
		end = trg[axys]
	}

	if axys == 0 {
		for i := start; i <= end; i++ {
			(*matrix)[src[1]][i]++
		}
	} else {
		for i := start; i <= end; i++ {
			(*matrix)[i][src[0]]++
		}

	}
}

func connectDiagonally(vent Vent, matrix *[1000][1000]int) {
	fmt.Printf("is Diagonal? %v\n", vent)
	sourceX := vent.source[0]
	targetX := vent.target[0]

	sourceY := vent.source[1]
	targetY := vent.target[1]

	// var startX, startY, endX, endY int
	// if sourceX > targetX {
	// 	startX = targetX
	// 	endX = sourceX
	// } else {
	// 	startX = sourceX
	// 	endX = targetX
	// }

	// if sourceY > targetY {
	// 	startY = targetY
	// 	endY = sourceY
	// } else {
	// 	startY = sourceY
	// 	endY = targetY
	// }

	// for startX <= endX && startY <= endY {
	// 	matrix[startX][startY]++
	// 	startX++
	// 	startY++
	// }

	var posX, posY int
	posX = sourceX
	posY = sourceY

	for posX != targetX && posY != targetY {
		matrix[posX][posY]++

		if posX > targetX {
			posX--
		} else {
			posX++
		}

		if posY > targetY {
			posY--
		} else {
			posY++
		}
	}

}

func findOverlap(matrix *[1000][1000]int) int {
	result := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if (*matrix)[i][j] > 1 {
				result++
			}
		}
	}
	return result
}
func main() {
	lines := utils.ReadFIleAsStringLines("input.txt")
	vents := parseInput(lines)
	matrix := [1000][1000]int{}
	for _, v := range vents {
		connectPoints(v, &matrix, true)
	}
	res := findOverlap(&matrix)
	fmt.Println(res)

}
