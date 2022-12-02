package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type RpsHand int // A rock paper scissors hand

const (
	Rock     RpsHand = 0
	Paper    RpsHand = 1
	Scissors RpsHand = 2
)

type RpsOutcome int

const (
	Win  RpsOutcome = 0
	Loss RpsOutcome = 1
	Draw RpsOutcome = 2
)

func (t RpsHand) verses(hand RpsHand) RpsOutcome {
	switch hand {
	case t:
		return Draw
	case (t + 1) % 3:
		return Loss
	default:
		return Win
	}
}

type RpsGame struct {
	playerA RpsHand
	playerB RpsHand
}

func Play(game RpsGame) (result RpsOutcome, points int) {
	result = game.playerA.verses(game.playerB)
	points += int(game.playerA) + 1
	switch result {
	case Win:
		points += 6
	case Loss:
		points += 0
	case Draw:
		points += 3
	}
	return
}

func IntoGuidePartA(input string) (game RpsGame, err error) {
	// we need our second person to be playerA for the game struct to work
	switch input[0] {
	case 'A':
		game.playerB = Rock
	case 'B':
		game.playerB = Paper
	case 'C':
		game.playerB = Scissors
	}

	switch input[2] {
	case 'X':
		game.playerA = Rock
	case 'Y':
		game.playerA = Paper
	case 'Z':
		game.playerA = Scissors
	}

	return
}

func partBGuideCreator() func(string) (RpsGame, error) {
	// we need our second person to be playerA for the game struct to work
	losingLUT := map[RpsHand]RpsHand{
		Rock:     Scissors,
		Scissors: Paper,
		Paper:    Rock,
	}

	winningLUT := map[RpsHand]RpsHand{
		Rock:     Paper,
		Paper:    Scissors,
		Scissors: Rock,
	}

	return func(input string) (game RpsGame, err error) {
		switch input[0] {
		case 'A':
			game.playerB = Rock
		case 'B':
			game.playerB = Paper
		case 'C':
			game.playerB = Scissors
		}

		switch input[2] {
		case 'X':
			game.playerA = losingLUT[game.playerB]
		case 'Y':
			game.playerA = game.playerB
		case 'Z':
			game.playerA = winningLUT[game.playerB]
		}

		return
	}
}

var IntoGuidePartB func(string) (RpsGame, error) = partBGuideCreator()

func calculateMaxScore(guide []RpsGame) int {
	totalScore := 0
	for i := range guide {
		_, points := Play(guide[i])
		totalScore += points
	}
	return totalScore
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)
	}

	// parse game
	lines := strings.Split(string(puzzleInput), "\n")
	roundsA := make([]RpsGame, 0, len(lines))
	roundsB := make([]RpsGame, 0, len(lines))

	for i := range lines {
		if len(lines[i]) > 0 {
			round, _ := IntoGuidePartA(lines[i])
			roundsA = append(roundsA, round)
			round, _ = IntoGuidePartB(lines[i])
			roundsB = append(roundsB, round)
		}
	}

	partOne := calculateMaxScore(roundsA)
	partTwo := calculateMaxScore(roundsB)

	fmt.Printf("Part One Solution: %v\n", partOne)
	fmt.Printf("Part One Solution: %v\n", partTwo)

}
