package main

import "testing"

func TestGetIncreaseDepth(t *testing.T) {

	var example = []int{198, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	result := getIncreaseDepth(example)
	if result != 7 {
		t.Error("resutl is not 6")
	}

}

func TestSumMeasuramentsByThree(t *testing.T) {

	var example = []int{199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}
	result := sumMeasuramentsByThree(example)
	if result[0] != 607 {
		t.Errorf("should be 607, but is %v\n", result[0])
	}
	if result[1] != 618 {
		t.Errorf("should be 618, but is %v\n", result[1])
	}
	if result[2] != 618 {
		t.Errorf("should be 618, but is %v\n", result[2])
	}
	if result[3] != 617 {
		t.Errorf("should be 617, but is %v\n", result[3])
	}
}
