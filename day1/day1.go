package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func readFIle(str string) []int {
	file, err := os.Open(str)
	check(err)
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		check(err)

		lines = append(lines, i)
	}
	return lines
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	lines := readFIle("input.txt")
	result := getIncreaseDepth(lines)
	fmt.Printf("First result is %v\n", result)

	result = getIncreaseDepth(sumMeasuramentsByThree(lines))
	fmt.Printf("Second result is %v\n", result)

}
