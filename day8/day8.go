package main

import (
	"fmt"
	"strings"

	"github.com/blackdev1l/aoc2021/utils"
)

type Entry struct {
	input  []string
	output []string
}


func main() {
	input := utils.ReadFIleAsStringLines("input.txt")

	parsed := parseInput(input)
	result := 0
	for _, v := range parsed {
		result += countOutput(v)
	}
	fmt.Println(result)
}

func parseInput(input []string) []Entry {
	entries := []Entry{}
	for _, i := range input {
		var entry Entry
		splittedString := strings.Split(i, "|")

		entry.input = strings.Split(splittedString[0], " ")
		entry.output = strings.Split(splittedString[1], " ")

		entries = append(entries, entry)

	}

	return entries

}

func countOutput(entry Entry) int {
	segments := getMappedSegments()
	result := 0
	for _, e := range entry.output {
		if segments[len(e)] != 0 {
			result++
		}
	}

	return result
}

func getMappedSegments() map[int]int {
	segments := make(map[int]int)
	segments[2] = 1
	segments[4] = 4
	segments[3] = 7
	segments[7] = 8

	return segments
}

func mapInput(input string, mappina *map[int]string) {

	if len(input) == 2 {
		(*mappina)[1] = input
	}

	if len(input) == 4 {
		(*mappina)[4] = input
	}

	if len(input) == 3 {
		(*mappina)[7] = input
	}

	if len(input) == 7 {
		(*mappina)[8] = input
	}

	if len(input) == 6 {
		if()

	}

}


func isSix(input string, mappina *map[string]int) bool {
	var result bool = false
	var count = 0
	for _, m := range mappina {
		for _, v := range string {
			if(v == m) {
				count++
			}
		}
	}

	return count = 1 
}