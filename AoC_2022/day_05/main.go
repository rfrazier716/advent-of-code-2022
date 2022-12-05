package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack = []rune

type CraneInstruction struct {
	Amount int
	From   int
	To     int
}

func ParseInput(lines []string) (stacks []Stack, instructions []CraneInstruction) {
	loadingStacks := true
	matcher, _ := regexp.Compile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			loadingStacks = false
			continue
		}
		if loadingStacks {
			// reverse the slice
			toAppend := []rune(lines[i])
			for i, j := 0, len(toAppend)-1; i < j; i, j = i+1, j-1 {
				toAppend[i], toAppend[j] = toAppend[j], toAppend[i]
			}
			stacks = append(stacks, toAppend)

		} else {
			match := matcher.FindStringSubmatch(lines[i])
			amount, _ := strconv.Atoi(match[1])
			from, _ := strconv.Atoi(match[2])
			to, _ := strconv.Atoi(match[3])
			instructions = append(instructions, CraneInstruction{
				amount,
				from,
				to,
			})
		}
	}

	return
}

func PuzzlePartOne(stacks []Stack, instructions []CraneInstruction, multiMove bool) string {
	for _, instruction := range instructions {
		from := &stacks[instruction.From-1]
		to := &stacks[instruction.To-1]

		// if we can move multiple at a time we just append the slice
		if multiMove {
			*to = append(*to, (*from)[len(*from)-instruction.Amount:]...)
		} else {
			for count := 0; count < instruction.Amount; count++ {
				*to = append(*to, (*from)[len(*from)-count-1])
			}
		}

		*from = (*from)[:len(*from)-instruction.Amount]
	}

	output := make([]rune, 0)
	for _, val := range stacks {
		output = append(output, val[len(val)-1])
	}
	return string(output)
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)

	}
	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]
	stacks, instructions := ParseInput(lines)
	ptOneStack := make([]Stack, len(stacks))
	copy(ptOneStack, stacks)
	ptTwoStack := make([]Stack, len(stacks))
	copy(ptTwoStack, stacks)

	fmt.Printf("Part One Solution: %v\n", PuzzlePartOne(ptOneStack, instructions, false))
	fmt.Printf("Part Two Solution: %v\n", PuzzlePartOne(ptTwoStack, instructions, true))
}
