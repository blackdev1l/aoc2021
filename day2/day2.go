package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/blackdev1l/aoc2021/utils"
)

type Position struct {
	horizontal, depth, aim int
}

func (p *Position) calculate() int {
	return p.depth * p.horizontal
}

func (p *Position) movePart1(list []string) {
	for _, line := range list {
		words := strings.Fields(line)
		direction := words[0]
		value, err := strconv.Atoi(words[1])
		utils.Check(err)

		switch direction {
		case "up":
			p.depth -= value
		case "down":
			p.depth += value
		case "forward":
			p.horizontal += value
		}
	}
}

func (p *Position) movePart2(list []string) {
	for _, line := range list {
		words := strings.Fields(line)
		direction := words[0]
		value, err := strconv.Atoi(words[1])
		utils.Check(err)

		switch direction {
		case "up":
			p.aim -= value
		case "down":
			p.aim += value
		case "forward":
			p.horizontal += value
			p.depth = p.depth + (p.aim * value)
		}
	}
}

func main() {
	p := Position{}
	lines := utils.ReadFIleAsStringLines("input.txt")
	p.movePart1(lines)
	result := p.calculate()
	fmt.Printf("result of day 2 part 1 is %v\n", result)
	p = Position{}
	p.movePart2(lines)
	result = p.calculate()

	fmt.Printf("result of day 2 part 2 is %v\n", result)
}
