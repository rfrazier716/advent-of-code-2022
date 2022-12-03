package main

import (
	"testing"
)

func TestScoreCalculation(t *testing.T) {
	testTable := []struct {
		game     GameRound
		expected int
	}{
		{NewRound(Rock, Rock), 4},
		{NewRound(Rock, Paper), 1},
		{NewRound(Rock, Scissors), 7},
	}
	for _, test := range testTable {
		if test.game.Score != test.expected {
			t.Errorf("test failed with input %v. Expected %v, got %v", test.game, test.expected, test.game.Score)
		}
	}

}
