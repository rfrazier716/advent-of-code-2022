package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type RpsHand int    // A rock paper scissors hand
type RpsOutcome int // The outcome, not boolean because you can draw

type Guide struct {
	// Maybe a bit excessive to make a guide struct instead of just pass the string
	OpponentAction rune
	playerAction   rune
}

const (
	Rock     RpsHand = 0
	Paper    RpsHand = 1
	Scissors RpsHand = 2
)

const (
	Win  RpsOutcome = 0
	Loss RpsOutcome = 1
	Draw RpsOutcome = 2
)

type GameRound struct {
	Hand     RpsHand
	Opponent RpsHand
	Outcome  RpsOutcome
	Score    int
}

type Player interface {
	// Player is an interface that can play a game round given a guide input
	Play(guide Guide) GameRound
}

// A decoder Func is a wrapper that implements the Player interface
// lets us write a function to decode the guide input and then play around based on that input
type DecoderFunc func(Guide) (RpsHand, RpsHand)

func (f DecoderFunc) Play(guide Guide) (round GameRound) {
	player, opponent := f(guide)
	return NewRound(player, opponent)
}

func NewRound(player RpsHand, opponent RpsHand) (round GameRound) {
	round.Hand = player
	round.Opponent = opponent

	round.Score = int(round.Hand) + 1 // initial score based on what you played
	switch opponent {
	case player:
		// draw
		round.Outcome = Draw
		round.Score += 3
	case (player + 1) % 3:
		// Loss
		round.Outcome = Loss
	case (player - 1) % 3, 2:
		// Victory
		round.Score += 6
		round.Outcome = Win
	}

	return
}

func PartOneDecoder(guide Guide) (player RpsHand, opponent RpsHand) {
	switch guide.OpponentAction {
	case 'A':
		opponent = Rock
	case 'B':
		opponent = Paper
	case 'C':
		opponent = Scissors
	}

	switch guide.playerAction {
	case 'X':
		player = Rock
	case 'Y':
		player = Paper
	case 'Z':
		player = Scissors
	}

	return
}

func PartTwoDecoder(guide Guide) (player RpsHand, opponent RpsHand) {
	switch guide.OpponentAction {
	case 'A':
		opponent = Rock
	case 'B':
		opponent = Paper
	case 'C':
		opponent = Scissors
	}

	switch guide.playerAction {
	case 'X':
		player = (opponent - 1) % 3
		// need to add looping behaviour to this
		if player < 0 {
			player = 2
		}

	case 'Y':
		player = opponent
	case 'Z':
		player = (opponent + 1) % 3
	}

	return
}

func CalculateScore(strategyGuide []Guide, playstyle Player) (score int) {
	for i := range strategyGuide {
		score += playstyle.Play(strategyGuide[i]).Score
	}
	return
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)
	}

	// parse game
	lines := strings.Split(string(puzzleInput), "\n")

	// Create our Strategy Guide
	strategyGuide := make([]Guide, len(lines)-1)
	for i := range strategyGuide {
		strategyGuide[i] = Guide{rune(lines[i][0]), rune(lines[i][2])}
	}

	partOne := CalculateScore(strategyGuide, DecoderFunc(PartOneDecoder))
	partTwo := CalculateScore(strategyGuide, DecoderFunc(PartTwoDecoder))

	fmt.Printf("Part One Solution: %v\n", partOne)
	fmt.Printf("Part One Solution: %v\n", partTwo)
}
