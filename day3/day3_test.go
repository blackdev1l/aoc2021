package main

import "testing"

var example = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestGetGammaRateToString(t *testing.T) {
	result := getGammaRateToString(example)
	if result != "10110" {
		t.Error("result is not 10110")
	}

}

func TestBinaryToDecimal(t *testing.T) {
	res := binaryToDecimal("10110")
	if res != 22 {
		t.Error("res is not 22")
	}
}

func TestFlipBinary(t *testing.T) {
	res := flipBinary(22)
	if res != 9 {
		t.Errorf("res is not 9 but is %v\n", res)
	}

}

func TestResult(t *testing.T) {
	gamma := binaryToDecimal(getGammaRateToString(example))
	epsilon := flipBinary(gamma)
	result := gamma * epsilon
	if result != 198 {
		t.Errorf("wrong!! res is 198 but yours is %v\n", result)
	}
}
