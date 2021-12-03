package main

import "testing"

var example = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2"}

func TestCalculate(t *testing.T) {
	p := Position{horizontal: 3, depth: 5}
	result := p.calculate()
	if result != 15 {
		t.Error("wrong calc")
	}

}

func TestMove1(t *testing.T) {
	p := Position{}
	p.movePart1(example)
	result := p.calculate()
	if p.calculate() != 150 {
		t.Errorf("calc should be 150 but is %v\n", result)
	}
}

func TestMove2(t *testing.T) {
	p := Position{}
	p.movePart2(example)
	result := p.calculate()
	if p.calculate() != 900 {
		t.Errorf("calc should be 900 but is %v\n", result)
	}
}
