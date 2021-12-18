package main

import (
	"fmt"
	"testing"

	"github.com/blackdev1l/aoc2021/utils"
)

func TestPart1(t *testing.T) {
	input := utils.MakeMatrix(utils.ReadFIleAsStringLines("example.txt"))
	var cavern = Cavern{synchronized: false}
	cavern.Init(input)
	cavern.Print()
	for i := 0; i < 100; i++ {
		cavern.BumpEnergy()
	}

	cavern.Print()

	if cavern.flashes != 1656 {
		t.Error("Error part 1")
	}

	fmt.Println(cavern.flashes)

}

func TestPart2(t *testing.T) {
	input := utils.MakeMatrix(utils.ReadFIleAsStringLines("example.txt"))
	var cavern = Cavern{synchronized: false}
	cavern.Init(input)
	cavern.Print()
	var step = 0
	for !cavern.synchronized {
		step++
		cavern.BumpEnergy()
	}

	fmt.Println(step)
	if step != 195 {
		t.Error("Error part 2")
	}

}
