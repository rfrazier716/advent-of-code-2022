package main

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	testInput := []string{
		"Sab",
		"cdE",
	}
	expected := [][]int{
		{0, 0, 1},
		{2, 3, 25},
	}
	start, end, parsed := parseInput(testInput)
	for r := range parsed {
		for c := range parsed[r] {
			if exp, act := expected[r][c], parsed[r][c]; exp != act {
				t.Errorf("Parser Failed. Height Mismatch at Coordiante (%v,%v). Expected %v, got %v", r, c, exp, act)
			}
		}
	}
	if exp, act := (coord2D{0, 0}), start; exp != act {
		t.Errorf("Parser Failed. Incorrect Start Coordinate, Expected %v, got %v", exp, act)
	}
	if exp, act := (coord2D{1, 2}), end; exp != act {
		t.Errorf("Parser Failed. Incorrect End Coordinate, Expected %v, got %v", exp, act)
	}
}

func TestDistanceCalculation(t *testing.T) {
	testInput := []string{
		"Swx",
		"Ezy",
	}
	expected := [][]int{
		{0, 4, 3},
		{0, 1, 2},
	}
	_, end, parsed := parseInput(testInput)
	minDistance := minDistanceToEnd(parsed, end)
	fmt.Println(minDistance)
	for r := range minDistance {
		for c := range minDistance[r] {
			if exp, act := expected[r][c], minDistance[r][c]; exp != act {
				t.Errorf("Parser Failed. Height Mismatch at Coordiante (%v,%v). Expected %v, got %v", r, c, exp, act)
			}
		}
	}
}

func TestPartOne(t *testing.T) {
	testInput := []string{
		"Sabqponm",
		"abcryxxl",
		"accszExk",
		"acctuvwj",
		"abdefghi",
	}
	start, end, heights := parseInput(testInput)
	fmt.Println(heights)
	if exp, act := 31, minDistanceToEnd(heights, start)[end.row][end.col]; exp != act {
		t.Errorf("Part One Failed. Incorrect Start Coordinate, Expected %v, got %v", exp, act)
	}

}
