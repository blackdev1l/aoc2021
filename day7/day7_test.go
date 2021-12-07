package main

import "testing"

var example = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func TestGestCost1(t *testing.T) {
	result := getCost(example, 2, false)
	if result != 37 {
		t.Errorf("Result is %v", result)
	}

}

func TestGestCost2(t *testing.T) {
	result := getCost(example, 2, true)
	if result != 206 {
		t.Errorf("Result is %v", result)
	}

}

func TestSumSteps(t *testing.T) {
	result := sumSteps(float64(4))
	if result != 10 {
		t.Errorf("result is %v\n", result)
	}
}
