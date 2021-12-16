package main

import (
	"fmt"
	"sort"
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

func getMatchingBrackets() map[string]string {
	var matchingBrackets = make(map[string]string)
	matchingBrackets["("] = ")"
	matchingBrackets["["] = "]"
	matchingBrackets["{"] = "}"
	matchingBrackets["<"] = ">"
	return matchingBrackets
}

func checkSyntax(line string) (int, Stack) {
	stack := Stack{}
	matchingBrackets := getMatchingBrackets()

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
				//fmt.Printf("Syntax Error expected %v but found %v\n", matchingBrackets[last], string(v))
				return points[string(v)], Stack{}
			}
		}
	}

	return 0, stack
}

func completeSyntax(line Stack) int {
	fmt.Printf("result for syntax %v", line.value)
	matchingBrackets := getMatchingBrackets()

	var points = make(map[string]int)
	points[")"] = 1
	points["]"] = 2
	points["}"] = 3
	points[">"] = 4

	var res = 0
	for len(line.value) != 0 {
		val := line.pop()
		res *= 5
		res += points[matchingBrackets[val]]

	}
	fmt.Printf("\t\t%v\n", res)
	return res
}

func getMiddleResult(scores []int) int {
	sort.Ints(scores)
	middle := (len(scores) / 2)
	return scores[middle]
}

func main() {
	lines := utils.ReadFIleAsStringLines("input.txt")
	var res = 0
	completedScores := []int{}
	for _, line := range lines {
		val, stack := checkSyntax(line)
		res += val
		if val == 0 {
			completedScores = append(completedScores, completeSyntax(stack))
		}

	}

	fmt.Println(res)
	fmt.Println(getMiddleResult(completedScores))

}
