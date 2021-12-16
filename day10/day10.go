package main

import (
	"fmt"
	"strings"

	"github.com/blackdev1l/aoc2021/utils"
)

type Stack struct {
	value []string
}

var openBrackets = "([{<"

func (s *Stack) push(val string) {
	s.value = append(s.value, val)
}

func (s *Stack) pop() string {
	val := s.value[len(s.value)-1]
	s.value = s.value[0 : len(s.value)-1]
	return val
}

func parseInput(filename string) []Stack {
	lines := utils.ReadFIleAsStringLines(filename)

	stacks := []Stack{}
	for _, line := range lines {
		stack := Stack{}
		for _, r := range line {
			stack.push(string(r))
		}
		stacks = append(stacks, stack)
	}

	return stacks
}

func checkSyntax(line string) int {
	stack := Stack{}
	var matchingBrackets = make(map[string]string)
	matchingBrackets["("] = ")"
	matchingBrackets["["] = "]"
	matchingBrackets["{"] = "}"
	matchingBrackets["<"] = ">"

	var points = make(map[string]int)
	points[")"] = 3
	points["]"] = 57
	points["}"] = 1197
	points[">"] = 25137

	for _, v := range line {
		if strings.Contains(openBrackets, string(v)) {
			stack.push(string(v))
		} else {
			last := stack.pop()
			if matchingBrackets[last] != string(v) {
				fmt.Printf("Syntax Error expected %v but found %v\n", matchingBrackets[last], string(v))
				return points[string(v)]
			}
		}
	}

	return 0

}

func main() {
	lines := utils.ReadFIleAsStringLines("input.txt")
	var res = 0
	for _, line := range lines {
		res += checkSyntax(line)
	}
	fmt.Println(res)
}
