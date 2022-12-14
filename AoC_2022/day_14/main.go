package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coord2D struct {
	row int
	col int
}

type SandTrap struct {
	sandSource   Coord2D
	crossSection map[Coord2D]struct{}
	sand         map[Coord2D]struct{}
	min          Coord2D // the top left corner of the cross section
	max          Coord2D // the bottom right corner of the cross section
}

func (st SandTrap) Display() {
	for r := st.min.row; r <= st.max.row; r++ {
		for c := st.min.col; c <= st.max.col; c++ {
			coord := Coord2D{r, c}
			toPrint := "."
			if _, ok := st.crossSection[coord]; ok {
				toPrint = "#"
			} else if _, ok := st.sand[coord]; ok {
				toPrint = "o"
			}
			fmt.Print(toPrint)
		}
		fmt.Println()
	}
}

func (st SandTrap) IsOccupied(coord Coord2D) bool {
	_, inCrossSection := st.crossSection[coord]
	_, inSand := st.sand[coord]
	return inCrossSection || inSand
}

func (st SandTrap) IsInBounds(coord Coord2D) bool {
	return coord.row >= st.min.row &&
		coord.row <= st.max.row &&
		coord.col >= st.min.col &&
		coord.col <= st.max.col
}

func (st *SandTrap) drop() (Coord2D, bool) {
	// create a new grain of sand
	// drop it until it exists the simulation bounds or is stuck
	prev, next := Coord2D{st.sandSource.row - 1, st.sandSource.col}, st.sandSource
	for prev != next && st.IsInBounds(next) {
		prev = next
		if down := (Coord2D{prev.row + 1, prev.col}); !st.IsOccupied(down) {
			next = down
			//fmt.Println("Down Works")
		} else if left := (Coord2D{prev.row + 1, prev.col - 1}); !st.IsOccupied(left) {
			//fmt.Println("Left Works")

			next = left
		} else if right := (Coord2D{prev.row + 1, prev.col + 1}); !st.IsOccupied(right) {
			//fmt.Println("Right Works")

			next = right
		}
	}
	// fmt.Println(prev, st.IsInBounds(prev))
	if st.IsInBounds(next) {
		st.sand[next] = struct{}{} // update the position in our cross section
		return next, true
	}

	return Coord2D{}, false
}

func (st *SandTrap) fill() int {
	// Fill drops sand until they start rolling into the abyss or fill the container
	// returns how much sand successfully dropped before the trap is filled
	i := 0

	for isFull := false; !isFull; {
		_, ok := st.drop()
		isFull = !ok
		if !isFull {
			i += 1
		}
		// fmt.Printf("\n----%v----\n",i)
		// st.Display()

	}

	return i
}

func ParseInput(input []string) SandTrap {
	matcher, _ := regexp.Compile(`([0-9]+),([0-9]+)`) // don't bother with errors for now

	// Constants for our SandTraps
	sandSource := Coord2D{0, 500} // Provided from Input
	min, max := sandSource, sandSource

	updateMinMax := func(row int, col int, currentMin *Coord2D, currentMax *Coord2D) {
		if row < currentMin.row {
			currentMin.row = row
		}
		if col < currentMin.col {
			currentMin.col = col
		}
		if row > currentMax.row {
			currentMax.row = row
		}
		if col > currentMax.col {
			currentMax.col = col
		}
	}

	// build the crossSection
	crossSection := make(map[Coord2D]struct{})
	for i := range input {
		matches := matcher.FindAllStringSubmatch(input[i], -1)

		// rows and columns are switch since they're doing x and y 
		rCur, _ := strconv.Atoi(matches[0][2])
		cCur, _ := strconv.Atoi(matches[0][1])
		for _, match := range matches[1:] {
			rNext, _ := strconv.Atoi(match[2])
			cNext, _ := strconv.Atoi(match[1])
			
			// swap orders if needed
			rMin, rMax := rCur, rNext
			if rCur > rNext {
				rMin, rMax = rNext, rCur
			}

			cMin, cMax := cCur, cNext
			if cCur > cNext {
				cMin, cMax = cNext, cCur
			}

			fmt.Println(rCur, rNext, cCur, cNext)
			for r := rMin; r <= rMax; r += 1 {
				for c := cMin; c <= cMax; c += 1 {
					crossSection[Coord2D{r, c}] = struct{}{}
					updateMinMax(r,c, &min, &max)
				}
				rCur, cCur = rNext, cNext
			}
		}
	}

	return SandTrap{
		sandSource: sandSource,
		crossSection: crossSection,
		sand: make(map[Coord2D]struct{}),
		min: min,
		max: max,
	}
}

func main() {
	puzzleInput, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Could not Load Puzzle Input: %v", err)
	}

	lines := strings.Split(string(puzzleInput), "\n")
	lines = lines[:len(lines)-1]
	
	trap := ParseInput(lines)
	
	fmt.Printf("Part One Solution: %v", trap.fill())
}
