package main

import (
	"fmt"
	"testing"
)

var example = []int{3, 4, 3, 1, 2}

func TestCount(t *testing.T) {
	fmt.Printf("Initial state: %v\n", example)
	for i := 1; i < 19; i++ {
		fmt.Printf("after %v days ", i)

		cycle(&example)
	}

}
