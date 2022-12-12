package main

import (
	"testing"
)

func TestRopeConstraint(t *testing.T) {
	testTable := []struct {
		rope         rope
		expectedTail coord2D
	}{
		{rope{&coord2D{2, 0}, &coord2D{0, 0}}, coord2D{1, 0}},
		{rope{&coord2D{2, 1}, &coord2D{0, 0}}, coord2D{1, 1}},
		{rope{&coord2D{-2, 0}, &coord2D{0, 0}}, coord2D{-1, 0}},
		{rope{&coord2D{-2, 1}, &coord2D{0, 0}}, coord2D{-1, 1}},
	}
	for i, test := range testTable {
		test.rope.constrain()
		if exp, act := test.expectedTail, *test.rope.tail; exp != act {
			t.Errorf("Test index %v failed. Expected %v, got %v", i, exp, act)
		}
	}
}

func TestPartOne(t *testing.T) {
	directions := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}
	rope := NewRope()
	if exp, act := 13, MovementTracker(directions, &rope); exp != act {
		t.Errorf("Part One failed, expected %v, got %v", exp, act)
	}
}

func TestPartTwo(t *testing.T) {
	directions := []string{
		"R 5",
		"U 8",
		"L 8",
		"D 3",
		"R 17",
		"D 10",
		"L 25",
		"U 20",
	}
	chain := NewRopeChain(9)
	if exp, act := 36, MovementTracker(directions, &chain); exp != act {
		t.Errorf("Part One failed, expected %v, got %v", exp, act)
	}
}
