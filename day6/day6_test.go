package main

import (
	"fmt"
	"testing"
)

var example = []int{3, 4, 3, 1, 2}

func TestCount(t *testing.T) {
	fmt.Printf("Initial state: %v\n", example)
	for i := 1; i < 257; i++ {
		fmt.Printf("after %v days \n", i)

		cycle(&example)
	}

}
