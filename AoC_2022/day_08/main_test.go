package main

import (
	"testing"
)

var testInput = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

var parsedTestInput = [][]int{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func TestStackAppend(t *testing.T) {
	trees := []*tree{
		{height: 1},
		{height: 2},
		{height: 3},
		{height: 2},
		{height: 3},
		{height: 4},
		{height: 1},
		{height: 5},
	}

	expectedStacks := []treeStack{
		{trees[:1], 0}, // 1
		{trees[:2], 1}, // 1 2
		{trees[:3], 2}, // 1 2 3
		{trees[:4], 2}, // 1 2 3 2
		{[]*tree{trees[0], trees[1], trees[2], trees[4]}, 2},           // 1 2 3 3
		{[]*tree{trees[0], trees[1], trees[2], trees[5]}, 3},           // 1 2 3 4
		{[]*tree{trees[0], trees[1], trees[2], trees[5], trees[6]}, 3}, // 1 2 3 4 1
		{[]*tree{trees[0], trees[1], trees[2], trees[5], trees[7]}, 4}, // 1 2 3 4 5
	}

	stack := newStack()
	for i, expected := range expectedStacks {
		updateStack(&stack, trees[i])
		if exp, act := expected.maxHeightIndex, stack.maxHeightIndex; exp != act {
			t.Errorf("Test failed on iteration %v, expected maxHeight index %v, got %v", i, exp, act)
		}
		if len(stack.trees) != len(expected.trees) {
			t.Errorf("Test failed on iteration %v, expected %v, got %v", i, expected, stack)
		} else {
			for j := range stack.trees {
				if stack.trees[j].height != expected.trees[j].height {
					t.Errorf("Stack mismatch at index %v, expected %v, got %v", j, expected.trees[j], stack.trees[j])
				}
			}
		}
	}
}

func TestParser(t *testing.T) {
	expected := parsedTestInput
	actual := parseInput(testInput)
	for r := range actual {
		for c := range actual[r] {
			if act, exp := actual[r][c], expected[r][c]; act != exp {
				t.Errorf("Error at index (%v, %v), expected %v, got %v", r, c, exp, act)
			}
		}
	}
}

func TestViewFromLeft(t *testing.T) {
	expected := [][]int{
		{0, 1, 2, 3, 1},
		{0, 1, 1, 1, 2},
		{0, 1, 1, 1, 1},
		{0, 1, 2, 1, 4},
		{0, 1, 1, 3, 1},
	}
	actual := furthestViewLookingLeft(parsedTestInput)
	for r := range actual {
		for c := range actual[r] {
			if act, exp := actual[r][c], expected[r][c]; act != exp {
				t.Errorf("Error at index (%v, %v), expected %v, got %v", r, c, exp, act)
			}
		}
	}
}

func TestViewFromRight(t *testing.T) {
	expected := [][]int{
		{2, 1, 1, 1, 0},
		{1, 1, 2, 1, 0},
		{4, 3, 1, 1, 0},
		{1, 1, 2, 1, 0},
		{1, 2, 1, 1, 0},
	}
	actual := furthestViewLookingRight(parsedTestInput)
	for r := range actual {
		for c := range actual[r] {
			if act, exp := actual[r][c], expected[r][c]; act != exp {
				t.Errorf("Error at index (%v, %v), expected %v, got %v", r, c, exp, act)
			}
		}
	}
}

func TestViewLookingUp(t *testing.T) {
	expected := [][]int{
		{0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1},
		{2, 1, 1, 2, 1},
		{1, 1, 2, 3, 3},
		{1, 2, 1, 4, 1},
	}
	actual := furthestViewLookingUp(parsedTestInput)
	for r := range actual {
		for c := range actual[r] {
			if act, exp := actual[r][c], expected[r][c]; act != exp {
				t.Errorf("Error at index (%v, %v), expected %v, got %v", r, c, exp, act)
			}
		}
	}
}

func TestViewLookingDown(t *testing.T) {
	expected := [][]int{
		{2, 1, 1, 4, 3},
		{1, 1, 2, 1, 1},
		{2, 2, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
	}
	actual := furthestViewLookingDown(parsedTestInput)
	for r := range actual {
		for c := range actual[r] {
			if act, exp := actual[r][c], expected[r][c]; act != exp {
				t.Errorf("Error at index (%v, %v), expected %v, got %v", r, c, exp, act)
			}
		}
	}
}

// func TestViews(t *testing.T) {
// 	views := calculateViews(parsedTestInput)
// 	testTable := []struct {
// 		row      int
// 		col      int
// 		visUp    bool
// 		visDown  bool
// 		visLeft  bool
// 		visRight bool
// 	}{
// 		{0, 0, true, false, true, false},
// 		{0, 0, true, false, false, true},
// 	}
// 	for i, testCase := range testTable {
// 		if act, exp := viewableFromLeft(parsedTestInput, views, testCase.row, testCase.col), testCase.visLeft; exp != act {
// 			t.Errorf("test failed at iteration %v, visible from left. Expected %v, got %v", i, exp, act)
// 		}
// 		if act, exp := viewableFromRight(parsedTestInput, views, testCase.row, testCase.col), testCase.visRight; exp != act {
// 			t.Errorf("test failed at iteration %v, visible from right. Expected %v, got %v", i, exp, act)
// 		}
// 		if act, exp := viewableFromAbove(parsedTestInput, views, testCase.row, testCase.col), testCase.visUp; exp != act {
// 			t.Errorf("test failed at iteration %v, visible from above. Expected %v, got %v", i, exp, act)
// 		}
// 		if act, exp := viewableFromBelow(parsedTestInput, views, testCase.row, testCase.col), testCase.visDown; exp != act {
// 			t.Errorf("test failed at iteration %v, visible from below. Expected %v, got %v", i, exp, act)
// 		}
// 	}
// }

func TestPartOne(t *testing.T) {
	expected := 21
	soln := PuzzlePartOne(testInput)
	if soln != 21 {
		t.Errorf("Part One Test failed, expected %v, got %v", expected, soln)
	}
}
