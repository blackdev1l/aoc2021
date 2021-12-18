package main

import (
	"fmt"

	"github.com/blackdev1l/aoc2021/utils"
)

type Octopus struct {
	hasFlashed bool
	energy     int
}

type Cavern struct {
	matrix       [][]Octopus
	flashes      int
	synchronized bool
}

func (c *Cavern) Init(input [][]int) {
	for i := 0; i < len(input); i++ {
		arr := []Octopus{}
		for j := 0; j < len(input[0]); j++ {
			oct := Octopus{energy: input[i][j], hasFlashed: false}
			arr = append(arr, oct)

		}
		c.matrix = append(c.matrix, arr)
	}
}

func (c *Cavern) BumpEnergy() {
	for i := 0; i < len(c.matrix); i++ {
		for j := 0; j < len(c.matrix[0]); j++ {
			c.matrix[i][j].energy++
		}
	}

	c.CheckForFlash()

	for i := 0; i < len(c.matrix); i++ {
		for j := 0; j < len(c.matrix[0]); j++ {
			oct := &c.matrix[i][j]
			if oct.energy > 9 {
				oct.hasFlashed = false
				oct.energy = 0
			}
		}
	}

	count := 0
	for i := 0; i < len(c.matrix); i++ {
		for j := 0; j < len(c.matrix[0]); j++ {
			oct := &c.matrix[i][j]
			if oct.energy == 0 {
				count++
			}
		}
	}

	if count == len(c.matrix[0])*len(c.matrix) {
		c.synchronized = true
	}

}

func (c *Cavern) CheckForFlash() {
	for i := 0; i < len(c.matrix); i++ {
		for j := 0; j < len(c.matrix[0]); j++ {
			oct := &c.matrix[i][j]
			if oct.energy > 9 && !oct.hasFlashed {
				oct.hasFlashed = true
				c.Flash(i, j)
				c.CheckForFlash()

			}
		}
	}
}

func (c *Cavern) Print() {
	for i := 0; i < len(c.matrix); i++ {
		for j := 0; j < len(c.matrix[0]); j++ {
			fmt.Printf("%v", c.matrix[i][j].energy)
		}
		fmt.Printf("\n")
	}
	fmt.Println(" --- ")
}

func (c *Cavern) Flash(i, j int) {
	c.flashes++

	// up
	if i != 0 {

		c.matrix[i-1][j].energy++

		// up left
		if j != 0 {
			c.matrix[i-1][j-1].energy++
		}

		// up right
		if j+1 < len(c.matrix[0]) {
			c.matrix[i-1][j+1].energy++
		}

	}

	// down
	if i+1 < len(c.matrix) {
		c.matrix[i+1][j].energy++

		// down left
		if j != 0 {
			c.matrix[i+1][j-1].energy++
		}

		// down right
		if j+1 < len(c.matrix[0]) {
			c.matrix[i+1][j+1].energy++
		}
	}

	//left
	if j != 0 {
		c.matrix[i][j-1].energy++
	}

	// right
	if j+1 < len(c.matrix[0]) {
		c.matrix[i][j+1].energy++
	}

}

func main() {
	input := utils.MakeMatrix(utils.ReadFIleAsStringLines("input.txt"))
	var cavern = Cavern{}
	cavern.Init(input)
	cavern.Print()
	for i := 0; i < 100; i++ {
		cavern.BumpEnergy()
	}

	cavern.Print()

	fmt.Println(cavern.flashes)

	cavern = Cavern{synchronized: false}
	cavern.Init(input)
	var step = 0
	for !cavern.synchronized {
		step++
		cavern.BumpEnergy()
	}

	fmt.Println(step)
}
