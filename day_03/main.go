package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Rucksack = string
type Priority = int

func FindPriority(sackItem rune) int {
	if sackItem > 'a' {
		return int(sackItem - 'a' + 1)
	} else {
		return int(sackItem - 'A' + 26 + 1)
	}
}

func FindDuplicateInPockets(rucksack Rucksack) Priority {
	// make an array to hold our counts
	seenLetters := make(map[rune]struct{})
	exists := struct{}{} // short hand
	for _, item := range rucksack[:len(rucksack)/2] {
		seenLetters[item] = exists
	}
	//fmt.Println(seenLetters)
	for _, item := range rucksack[len(rucksack)/2:] {
		if _, found := seenLetters[item]; found {
			return FindPriority(item)
			//fmt.Printf("Found Duplicate Item: %v", string(item))
		}
	}
	log.Printf("Could not find duplicate entry in rucksack: %v", rucksack)
	return 0 // should never get here
}

func FindCommonItem(packs []Rucksack) rune {
	itemCounter := make(map[rune]int)
	for i := range packs {
		for _, item := range packs[i] {
			if itemCounter[item] == i {
				// if we've seen it and it's appeared i times, add it to the counter
				itemCounter[item] = i + 1
			}
		}
	}
	for key, val := range itemCounter {
		if val == len(packs) {
			return key
		}
	}

	log.Printf("Could not find common item in Packs %v", packs)
	return '0'
}

func PartOne(packs []Rucksack) int {
	prioritySum := 0
	for i := range packs {
		prioritySum += FindDuplicateInPockets(packs[i])
	}
	return prioritySum
}

func PartTwo(packs []Rucksack) int {
	badgeSum := 0
	for i := 0; i < len(packs)/3; i++ {
		startIndex := 3 * i
		stopIndex := 3 * (i + 1)
		badgeSum += FindPriority(FindCommonItem(packs[startIndex:stopIndex]))
	}
	return badgeSum
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)

	}

	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]
	fmt.Println(PartOne(lines))
	fmt.Println(PartTwo(lines))

}
