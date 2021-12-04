package main

import (
	"fmt"

	"github.com/blackdev1l/aoc2021/utils"
)

func getIncreaseDepth(list []int) int {

	var count int
	var currentNumber int = list[0]
	for _, measurement := range list {
		if currentNumber < measurement {
			count++
		}
		currentNumber = measurement
	}

	return count
}

func sumMeasuramentsByThree(list []int) []int {
	var sumList []int
	for len(list) >= 3 {
		var count int
		for _, n := range list[:3] {
			count += n
		}
		sumList = append(sumList, count)
		list = list[1:]

	}
	return sumList
}

func main() {

	lines := utils.ReadFIleAsIntLines("input.txt")
	result := getIncreaseDepth(lines)
	fmt.Printf("First result is %v\n", result)

	result = getIncreaseDepth(sumMeasuramentsByThree(lines))
	fmt.Printf("Second result is %v\n", result)

}
