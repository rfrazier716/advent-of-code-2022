package main

import (
	"fmt"
	"log"
	"os"
)

func FindFirstUniqueSubstring(input []rune, len int) int {
	// returns the index of the first unique substring of the input string
	ItemsInWindow := make(map[rune]struct{})
	head, tail := 0,0
	for ;head - tail < len; head += 1{
		for _, ok := ItemsInWindow[input[head]]; ok; _, ok = ItemsInWindow[input[head]] {
			delete(ItemsInWindow, input[tail])
			tail += 1
		}
		// add our head to the map
		ItemsInWindow[input[head]] = struct{}{} // shorthand for a map
	}
	return head
}

func main() {
	inputFile, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)
	}

	puzzleInput := []rune(string(inputFile))

	fmt.Printf("Part One Solution: %v\n", FindFirstUniqueSubstring(puzzleInput, 4))
	fmt.Printf("Part Two Solution: %v\n", FindFirstUniqueSubstring(puzzleInput, 14))
}
