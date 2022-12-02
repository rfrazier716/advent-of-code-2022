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
	testTable := []struct {
		input    string
		expected RpsGame
	}{
		{"A Y", RpsGame{Rock, Rock}},
		{"B X", RpsGame{Rock, Paper}},
		{"C Z", RpsGame{Rock, Scissors}},
	}
	for _, test := range testTable {
		if actual, _ := IntoGuidePartB(test.input); test.expected != actual {
			t.Errorf("test failed with input %v. Expected %v, got %v", test.input, test.expected, actual)
		}
	}

}
