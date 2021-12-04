package utils

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFIleAsStringLines(str string) []string {
	file, err := os.Open(str)
	Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		Check(err)
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadFIleAsIntLines(str string) []int {
	file, err := os.Open(str)
	Check(err)
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		Check(err)

		lines = append(lines, i)
	}
	return lines
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
