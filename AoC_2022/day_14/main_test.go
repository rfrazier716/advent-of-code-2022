package main

import (
	"fmt"
	"testing"
)

func TestSomethingSilly(t *testing.T) {
	testTable := []struct {
		a        int
		b        int
		expected int
	}{{1, 2, 3}, {2, 3, 5}, {-4, 7, 3}}
	for i, test := range testTable {
		if res := test.a + test.b; res != test.expected {
			t.Errorf("Test index %v failed. Expected %v, got %v", i, test.expected, res)
		}
	}
}

func newTestTrap() SandTrap {
	testCaseCrossSections := []Coord2D{
		{4, 498},
		{5, 498},
		{6, 498},
		{6, 496},
		{6, 497},
		{9, 494},
		{9, 495},
		{9, 496},
		{9, 497},
		{9, 498},
		{9, 499},
		{9, 500},
		{9, 501},
		{9, 502},
		{8, 502},
		{7, 502},
		{6, 502},
		{5, 502},
		{4, 502},
		{4, 503},
	}
	sandSource := Coord2D{0, 500}
	min, max := sandSource, sandSource
	crossSection := make(map[Coord2D]struct{})
	for i, val := range testCaseCrossSections {
		// update min-max
		if val.row < min.row {
			min.row = val.row
		}
		if val.col < min.col {
			min.col = val.col
		}
		if val.row > max.row {
			max.row = val.row
		}
		if val.col > max.col {
			max.col = val.col
		}

		// add to struct
		crossSection[testCaseCrossSections[i]] = struct{}{}
	}

	return SandTrap{
		sandSource:   sandSource,
		crossSection: crossSection,
		sand:         make(map[Coord2D]struct{}),
		min:          min,
		max:          max,
	}
}

func TestSandDrop(t *testing.T) {
	trap := newTestTrap()
	fmt.Println(trap)
	expectedPositions := []Coord2D{
		{8, 500},
		{8, 499},
		{8, 501},
		{7, 500},
		{8, 498},
	}

	for i, act := range expectedPositions {
		exp, _ := trap.drop() // drop the next sand
		if exp != act {
			t.Errorf("Failed on iteration %v, expected %v, got %v", i, exp, act)
		}
	}
}

func TestSandFill(t *testing.T) {
	trap := newTestTrap()
	if exp, act := 24, trap.fill(); exp != act {
		t.Errorf("Part One Test Failed. Expected %v, got %v", exp, act)
	}
}
