package main

import (
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

func TestParser(t *testing.T) {
	input := []string{
		"ABC",
		"DEF",
		"GHI",
		"",
		"move 1 from 2 to 3",
		"move 2 from 3 to 4",
	}
	expectedStacks := []Stack{[]rune("CBA"), []rune("FED"), []rune("IHG")}
	expectedInstructions := []CraneInstruction{
		{1, 2, 3},
		{2, 3, 4},
	}

	actStacks, actInstructions := ParseInput(input)
	if len(expectedStacks) != len(actStacks) {
		t.Errorf("Parsed Stacks Incorrectly, expected %v, got %v", expectedStacks, actInstructions)
	} else {
		for i := range expectedStacks {
			if expectedStacks[i] != actStacks[i] {
				t.Errorf("Stack %v does not match expected. Expected %v, got %v", i, expectedStacks[i], actStacks[i])
			}
		}
	}
	if len(expectedInstructions) != len(actInstructions) {
		t.Errorf("Parsed Instructions Incorrectly, expected %v, got %v", expectedInstructions, actInstructions)
	} else {
		for i := range expectedInstructions {
			if expectedInstructions[i] != actInstructions[i] {
				t.Errorf("Instruction %v does not match expected. Expected %v, got %v", i, expectedInstructions[i], actInstructions[i])
			}
		}
	}
}
