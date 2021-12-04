package main

import (
	"fmt"
	"strconv"

	"github.com/blackdev1l/aoc2021/utils"
)

func getGammaRateToString(list []string) string {
	sizeOfLine := len(list[0])
	numberOfLines := len(list)

	var result string

	for i := 0; i < sizeOfLine; i++ {
		counter := make(map[string]int)
		for j := 0; j < numberOfLines; j++ {
			line := []rune(list[j])
			char := string(line[i])
			counter[char] += 1
		}
		if counter["0"] > counter["1"] {
			result += "0"
		} else {
			result += "1"
		}
	}

	return result
}

func binaryToDecimal(input string) int {
	res, err := strconv.ParseInt(input, 2, 64)
	utils.Check(err)

	fmt.Println(res)
	return int(res)
}

func flipBinary(val string) int {
	var temp string

	for _, char := range val {
		char, err := strconv.Atoi(string(char))
		utils.Check(err)
		if char == 0 {
			temp += "1"
		} else {
			temp += "0"
		}
	}
	return binaryToDecimal(temp)
}

func main() {
	list := utils.ReadFIleAsStringLines("input.txt")
	binString := getGammaRateToString(list)
	gamma := binaryToDecimal(binString)
	epsilon := flipBinary(binString)
	fmt.Printf("result is %v\n", gamma*epsilon)
}
