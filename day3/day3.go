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

func oxygenGeneratorRating(list []string, offset int) string {
	fmt.Printf("offset is %v\n", offset)
	if len(list) == 1 {
		return list[0]
	}
	numberOfLines := len(list)

	var result []string

	counter := make(map[string]int)
	resolv := make(map[int][]string)
	for j := 0; j < numberOfLines; j++ {
		line := []rune(list[j])
		char := string(line[offset])
		counter[char] += 1
		if line[offset] == '0' {
			resolv[0] = append(resolv[0], list[j])
		} else {
			resolv[1] = append(resolv[1], list[j])
		}
	}
	if counter["1"] >= counter["0"] {
		result = resolv[1]

	} else {
		result = resolv[0]
	}
	fmt.Printf("counter of 1 is %v counter of 0 is %v \n", counter["1"], counter["0"])

	fmt.Println(result)

	return oxygenGeneratorRating(result, offset+1)
}

func co2Scrubber(list []string, offset int) string {
	fmt.Printf("offset is %v\n", offset)
	if len(list) == 1 {
		return list[0]
	}
	numberOfLines := len(list)

	var result []string

	counter := make(map[string]int)
	resolv := make(map[int][]string)
	for j := 0; j < numberOfLines; j++ {
		line := []rune(list[j])
		char := string(line[offset])
		counter[char] += 1
		if line[offset] == '0' {
			resolv[0] = append(resolv[0], list[j])
		} else {
			resolv[1] = append(resolv[1], list[j])
		}
	}
	if counter["1"] < counter["0"] {
		result = resolv[1]
	} else {
		result = resolv[0]
	}
	fmt.Printf("counter of 1 is %v counter of 0 is %v \n", counter["1"], counter["0"])

	fmt.Println(result)

	return co2Scrubber(result, offset+1)
}

func main() {
	list := utils.ReadFIleAsStringLines("input.txt")
	binString := getGammaRateToString(list)
	gamma := binaryToDecimal(binString)
	epsilon := flipBinary(binString)
	fmt.Printf("result part 1 is %v\n", gamma*epsilon)

	co2 := binaryToDecimal(co2Scrubber(list, 0))
	oxygen := binaryToDecimal(oxygenGeneratorRating(list, 0))
	fmt.Printf("result part 2 is %v\n", co2*oxygen)
}
