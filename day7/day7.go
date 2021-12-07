package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/blackdev1l/aoc2021/utils"
)

func getCost(list []int, position int, part2 bool) int {
	result := 0
	for _, i := range list {
		tmp := math.Abs(float64(i - position))
		if part2 {
			tmp = float64(sumSteps(tmp))
		}
		result += int(tmp)

	}
	fmt.Printf("cost is %v\n", result)

	return result
}

func parseInput(list string) []int {
	var result []int
	tmp := strings.Split(list, ",")
	for _, i := range tmp {
		value, err := strconv.Atoi(i)
		utils.Check(err)
		result = append(result, value)
	}
	return result
}
func sumSteps(count float64) int {
	result := 0
	for i := 1; i <= int(count); i++ {
		result += i

	}
	return result
}

func main() {
	file := utils.ReadFIleAsStringLines("input.txt")
	list := parseInput(file[0])
	result := -1
	for _, i := range list {
		cost := getCost(list, i, true)
		if cost < result || result == -1 {
			result = cost
		}

	}

	fmt.Println(result)

}
