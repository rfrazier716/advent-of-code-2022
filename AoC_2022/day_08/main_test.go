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

func TestPartOne(t *testing.T) {
	expected := 21
	soln := PuzzlePartOne(testInput)
	if soln != 21 {
		t.Errorf("Part One Test failed, expected %v, got %v", expected, soln)
	}
}
