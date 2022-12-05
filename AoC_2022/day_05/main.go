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

func clone(stacks []Stack) []Stack {
	// Deep-copy a Stack
	// Since it's a slice of slices, we cannot use teh built in copy method
	
	cloned := make([]Stack, len(stacks))
	for i := range stacks {
		cloned[i] = make(Stack, len(stacks[i]))
		copy(cloned[i], stacks[i])
	}

	return cloned
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

func Process(stacks []Stack, instructions []CraneInstruction, multiMove bool) []Stack {
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

	return stacks
}

func GetUppermostValues(stacks []Stack) string {
	// Gets the top values of the stack
	output := make([]rune, 0)
	for _, val := range stacks {
		output = append(output, val[len(val)-1])
	}
	return string(output)
}

func PuzzlePartOne(stacks []Stack, instructions []CraneInstruction) string {
	// process the instructions
	processed := Process(stacks, instructions, false)
	return GetUppermostValues(processed)
}


func PuzzlePartTwo(stacks []Stack, instructions []CraneInstruction) string {
	// process the instructions
	processed := Process(stacks, instructions, true)
	return GetUppermostValues(processed)
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)

	}
	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]
	stacks, instructions := ParseInput(lines)

	// need to clone our stacks since the function mutates it
	fmt.Printf("Part One Solution: %v\n", PuzzlePartOne(clone(stacks), instructions))
	fmt.Printf("Part Two Solution: %v\n", PuzzlePartTwo(stacks, instructions))
}
