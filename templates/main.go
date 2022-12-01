package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)

	}

	fmt.Print(puzzleInput)
}
