package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CleaningRange struct {
	Start int
	End   int
}

type ElfPair struct {
	First  CleaningRange
	Second CleaningRange
}

func IsFullyRedundant(pair ElfPair) bool {
	firstEnclosesSecond := pair.First.Start <= pair.Second.Start && pair.First.End >= pair.Second.End
	secondEnclosesFirst := pair.Second.Start <= pair.First.Start && pair.Second.End >= pair.First.End
	return firstEnclosesSecond || secondEnclosesFirst
}

func IsPartiallyRedundant(pair ElfPair) bool {
	if pair.First.Start <= pair.Second.Start {
		return pair.First.End >= pair.Second.Start
	} else {
		return pair.Second.End >= pair.First.Start
	}
}

func ParseInput(lines []string) []ElfPair {
	matcher, _ := regexp.Compile(`([0-9]+)-([0-9]+)`)

	pairs := make([]ElfPair, len(lines))
	for i := range lines {
		matches := matcher.FindAllStringSubmatch(lines[i], -1)
		elfOneStart, _ := strconv.Atoi(matches[0][1])
		elfOneEnd, _ := strconv.Atoi(matches[0][2])
		elfTwoStart, _ := strconv.Atoi(matches[1][1])
		elfTwoEnd, _ := strconv.Atoi(matches[1][2])

		pairs[i] = ElfPair{
			CleaningRange{elfOneStart, elfOneEnd},
			CleaningRange{elfTwoStart, elfTwoEnd},
		}
	}

	return pairs
}

func Solver(pairings []ElfPair) (int, int) {
	partOneCounter, partTwoCounter := 0, 0
	for i := range pairings {
		if IsFullyRedundant(pairings[i]) {
			partOneCounter += 1
		}
		if IsPartiallyRedundant(pairings[i]) {
			partTwoCounter += 1
		}
	}
	return partOneCounter, partTwoCounter
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)
	}
	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]
	pairs := ParseInput(lines)

	partOneSoln, partTwoSoln := Solver(pairs)

	fmt.Printf("Part One: %v\n", partOneSoln)
	fmt.Printf("Part Two: %v\n", partTwoSoln)

}
