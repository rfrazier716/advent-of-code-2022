package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(lines []string) []int {
	regValue := make([]int, 0)
	prev := 1
	for i := range lines {
		valToAdd := 0
		var toAppend []int
		switch lines[i][:4] {
		case "noop":
			toAppend = []int{prev}
		case "addx":
			valToAdd, _ = strconv.Atoi(lines[i][5:])
			toAppend = []int{prev, prev}
		}
		regValue = append(regValue, toAppend...)
		prev += valToAdd
	}
	return regValue
}

func signalStrengthSum(signal []int, cycles ...int) int {
	sum := 0
	for _, val := range cycles {
		sum += signal[val-1] * val
	}
	return sum
}

func display(register []int) {
	for i := range register {
		if i%40 == 0 {
			fmt.Println()
		}
		if register[i] >= (i%40)-1 && register[i] <= (i%40)+1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
	}
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)

	}

	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]
	registers := parseInput(lines)
	partOne := signalStrengthSum(registers, 20, 60, 100, 140, 180, 220)
	fmt.Println(partOne)
	display(registers)
}
